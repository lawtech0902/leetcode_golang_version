/*
__author__ = 'lawtech'
__date__ = '2020/03/11 1:59 下午'
*/

package _152

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minTmp, maxTmp, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		a := nums[i] * maxTmp
		b := nums[i] * minTmp
		c := nums[i]
		maxTmp = max(a, b, c)
		minTmp = min(a, b, c)
		if maxTmp > res {
			res = maxTmp
		}
	}

	return res
}

func max(nums ...int) int {
	res := math.MinInt64
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

func min(nums ...int) int {
	res := math.MaxInt64
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}
