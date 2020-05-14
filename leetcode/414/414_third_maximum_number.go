/*
__author__ = 'lawtech'
__date__ = '2020/05/14 10:04 上午'
*/

package _414

import "math"

// 设置一大，二大，三大 O(n) O(1)
func thirdMax(nums []int) int {
	max, secondMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > max {
			thirdMax = secondMax
			secondMax = max
			max = num
		} else if num > secondMax && num < max {
			thirdMax = secondMax
			secondMax = num
		} else if num > thirdMax && num < secondMax {
			thirdMax = num
		}
	}

	if thirdMax == math.MinInt64 {
		return max
	}

	return thirdMax
}
