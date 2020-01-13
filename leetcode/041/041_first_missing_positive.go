/*
__author__ = 'lawtech'
__date__ = '2020/1/13 4:22 下午'
*/

package _41

import "math"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	contains := 0
	for _, num := range nums {
		if num == 1 {
			contains++
			break
		}
	}

	// 1 not in nums
	if contains == 0 {
		return 1
	}

	// nums == [1]
	if n == 1 {
		return 2
	}

	// 用1替换非正数和大于n的数
	// 转换之后nums只包含正数
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	for i := 0; i < n; i++ {
		a := int(math.Abs(float64(nums[i])))
		if a == n {
			nums[0] = -int(math.Abs(float64(nums[0])))
		} else {
			nums[a] = -int(math.Abs(float64(nums[a])))
		}
	}

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return n
	}

	return n + 1
}
