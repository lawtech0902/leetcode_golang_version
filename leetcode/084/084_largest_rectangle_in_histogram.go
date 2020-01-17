/*
__author__ = 'lawtech'
__date__ = '2020/1/17 2:45 下午'
*/

package _84

import "math"

func largestRectangleArea(heights []int) int {
	var stack []int
	i := 0
	res := 0

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			res = int(math.Max(float64(res), float64(width*heights[cur])))
			i -= 1
		}

		i += 1
	}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		var width int
		if len(stack) == 0 {
			width = i
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}

		res = int(math.Max(float64(res), float64(width*heights[cur])))
	}

	return res
}
