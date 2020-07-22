/*
__author__ = 'lawtech'
__date__ = '2020/07/22 2:58 下午'
*/

package _763

import "math"

func partitionLabels(S string) []int {
	last := make([]int, 26)
	for i, ch := range S {
		last[ch-'a'] = i
	}

	start, end := 0, 0
	var res []int
	for i := 0; i < len(S); i++ {
		end = max(end, last[S[i]-'a'])
		if i == end {
			res = append(res, i-start+1)
			start = i + 1
		}
	}

	return res
}

func max(values ...int) int {
	maxValue := math.MinInt32
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}
