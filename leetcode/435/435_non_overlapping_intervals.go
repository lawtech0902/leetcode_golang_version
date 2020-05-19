/*
__author__ = 'lawtech'
__date__ = '2020/05/19 10:22 上午'
*/

package _435

import "sort"

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (iv Intervals) Len() int {
	return len(iv)
}

func (iv Intervals) Less(i, j int) bool {
	return iv[i].End < iv[j].End
}

func (iv Intervals) Swap(i, j int) {
	iv[i], iv[j] = iv[j], iv[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	var ivs Intervals
	for _, interval := range intervals {
		ivs = append(ivs, Interval{
			Start: interval[0],
			End:   interval[1],
		})
	}

	sort.Sort(ivs)

	count := 1
	start := ivs[0].End
	for i := 1; i < len(intervals); i++ {
		if ivs[i].Start < start {
			continue
		}

		count++
		start = ivs[i].End
	}

	return len(ivs) - count
}
