/*
__author__ = 'lawtech'
__date__ = '2020/05/22 10:18 上午'
*/

package _453

import "math"

func minMoves(nums []int) int {
	res, minVal := 0, math.MaxInt64
	for _, num := range nums {
		minVal = min(minVal, num)
	}

	for _, num := range nums {
		res += num - minVal
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
