/*
__author__ = 'lawtech'
__date__ = '2020/1/16 10:11 上午'
*/

package _56

import (
	"math"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
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
	temp := its[0]
	for i := 1; i < len(its); i++ {
		if its[i].Start <= temp.End {
			temp.End = int(math.Max(float64(temp.End), float64(its[i].End)))
		} else {
			newIts = append(newIts, temp)
			temp = its[i]
		}
	}

	newIts = append(newIts, temp)
	var res [][]int
	for _, newIt := range newIts {
		res = append(res, []int{newIt.Start, newIt.End})
	}
	return res
}
