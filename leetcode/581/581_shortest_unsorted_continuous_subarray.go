/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-22 09:58:00
 */

package _581

import (
	"container/list"
	"math"
	"sort"
)

// 排序比较
func findUnsortedSubarray(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	l, r := 0, -1
	for i := range sorted {
		if sorted[i] != nums[i] {
			l = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if sorted[i] != nums[i] {
			r = i
			break
		}
	}

	return r - l + 1
}

// 栈
func findUnsortedSubarray1(nums []int) int {
	stack := list.New()
	l, r := len(nums), 0
	for i := 0; i < len(nums); i++ {
		for stack.Len() != 0 && nums[stack.Back().Value.(int)] > nums[i] {
			l = int(math.Min(float64(l), float64(stack.Back().Value.(int))))
			stack.Remove(stack.Back())
		}

		stack.PushBack(i)
	}

	stack.Init()
	for i := len(nums) - 1; i >= 0; i-- {
		for stack.Len() != 0 && nums[stack.Back().Value.(int)] < nums[i] {
			r = int(math.Max(float64(r), float64(stack.Back().Value.(int))))
			stack.Remove(stack.Back())
		}
	}

	if r > l {
		return r - l + 1
	}

	return 0
}

// 不需要额外空间
func findUnsortedSubarray2(nums []int) int {
	start, end, min, max := 0, -1, math.MaxInt32, math.MinInt32
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}

		if nums[j] <= min {
			min = nums[j]
		} else {
			start = j
		}
	}

	return end - start + 1
}
