/*
__author__ = 'lawtech'
__date__ = '2020/07/23 11:04 上午'
*/

package _769

import "math"

// 暴力
func maxChunksToSorted(arr []int) int {
	res, maxVal := 0, 0
	for i := 0; i < len(arr); i++ {
		maxVal = max(maxVal, arr[i])
		if maxVal == i {
			res++
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
