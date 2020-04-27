/*
__author__ = 'lawtech'
__date__ = '2020/04/27 4:29 下午'
*/

package _334

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	low, medium := nums[0], math.MaxInt32
	for _, v := range nums {
		if v > medium {
			return true
		}

		if v < medium && v > low {
			medium = v
		}

		if v < low {
			low = v
		}
	}

	return false
}
