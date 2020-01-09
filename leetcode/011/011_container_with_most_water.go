/*
__author__ = 'lawtech'
__date__ = '2020/1/9 4:36 下午'
*/

package _11

import "math"

func maxArea(height []int) int {
	// 双指针法
	left, right := 0, len(height)-1
	res, area := 0, 0

	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left += 1
		} else {
			area = height[right] * (right - left)
			right -= 1
		}
		res = int(math.Max(float64(res), float64(area)))
	}

	return res
}
