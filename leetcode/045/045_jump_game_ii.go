/*
__author__ = 'lawtech'
__date__ = '2020/1/14 1:35 下午'
*/

package _45

import "math"

func jump(nums []int) int {
	end, maxPosition, steps := 0, 0, 0
	size := len(nums)

	for i := 0; i < size-1; i++ {
		// 找能跳最远的
		maxPosition = int(math.Max(float64(maxPosition), float64(nums[i]+i)))
		if i == end && end < size-1 { // 遇到边界，就更新边界，并且步数加一
			end = maxPosition
			steps += 1
		}
	}

	return steps
}
