package _53

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午11:11'
*/

import "math"

func maxSubArray(nums []int) int {
	maxSum, preSum := nums[0], 0

	for _, num := range nums {
		if preSum <= 0 {
			preSum = num
		} else {
			preSum += num
		}

		maxSum = int(math.Max(float64(maxSum), float64(preSum)))
	}

	return maxSum
}
