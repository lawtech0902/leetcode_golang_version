/*
__author__ = 'lawtech'
__date__ = '2020/1/16 10:51 上午'
*/

package _57

import (
	"math"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	size1 := len(intervals)
	if size1 == 1 {
		return intervals
	}

	its := make([]Interval, 0)
	for _, it := range intervals {
		its = append(its, Interval{it[0], it[1]})
	}

	sort.Slice(its, func(i, j int) bool {
		return its[i].Start < its[j].Start
	})

	var newIts []Interval
	for i := 0; i < size1; i++ {
		if len(newIts) == 0 {
			newIts = append(newIts, its[i])
		} else {
			size2 := len(newIts)
			if newIts[size2-1].Start <= its[i].Start && its[i].Start <= newIts[size2-1].End {
				newIts[size2-1].End = int(math.Max(float64(its[i].End), float64(newIts[size2-1].End)))
			} else {
				newIts = append(newIts, its[i])
			}
		}
	}

	var res [][]int
	for _, newIt := range newIts {
		res = append(res, []int{newIt.Start, newIt.End})
	}
	return res
}
