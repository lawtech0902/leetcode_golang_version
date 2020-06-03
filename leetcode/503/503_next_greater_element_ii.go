/*
__author__ = 'lawtech'
__date__ = '2020/06/03 11:21 上午'
*/

package _503

import "container/list"

// 单调栈
func nextGreaterElements(nums []int) []int {
	var stack list.List
	res := make([]int, len(nums))

	for i := 2*len(nums) - 1; i >= 0; i-- {
		for stack.Len() != 0 && nums[stack.Back().Value.(int)] <= nums[i%len(nums)] {
			stack.Remove(stack.Back())
		}

		if stack.Len() == 0 {
			res[i%len(nums)] = -1
		} else {
			res[i%len(nums)] = nums[stack.Back().Value.(int)]
		}

		stack.PushBack(i % len(nums))
	}

	return res
}
