/*
__author__ = 'lawtech'
__date__ = '2020/1/13 4:47 下午'
*/

package _42

func trap(height []int) int {
	res, size := 0, len(height)
	if size == 0 {
		return 0
	}

	leftMax, rightMax := 0, 0
	left, right := 0, size-1

	for left < right {
		if height[left] < height[right] {
			if height[left] < leftMax {
				res += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left += 1
		} else {
			if height[right] < rightMax {
				res += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right -= 1
		}
	}

	return res
}
