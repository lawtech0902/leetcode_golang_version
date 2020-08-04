/*
__author__ = 'lawtech'
__date__ = '2020/08/04 11:08 上午'
*/

package _42

import "math"

func maxSubArray(nums []int) int {
	maxSum, preSum := nums[0], 0
	for _, num := range nums {
		if preSum <= 0 {
			preSum = num
		} else {
			preSum += num
		}

		maxSum = max(maxSum, preSum)
	}

	return maxSum
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
