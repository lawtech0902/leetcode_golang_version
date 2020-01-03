package _53

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午11:11'
*/

func maxSubArray(nums []int) int {
	max_sum := nums[0]
	pre_sum := 0

	for _, num := range nums {
		if pre_sum < 0 {
			pre_sum = num
		} else {
			pre_sum += num
		}

		max_sum = int(math.Max(float64(max_sum), float64(pre_sum)))
	}

	return max_sum
}
