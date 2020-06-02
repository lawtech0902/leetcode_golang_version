/*
__author__ = 'lawtech'
__date__ = '2020/06/02 2:12 下午'
*/

package _496

import "container/list"

// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var (
		res   []int
		stack list.List
	)

	hash := make(map[int]int)
	for _, v := range nums2 {
		for stack.Len() > 0 && stack.Back().Value.(int) < v {
			hash[stack.Remove(stack.Back()).(int)] = v
		}

		stack.PushBack(v)
	}

	for _, v := range nums1 {
		if _, ok := hash[v]; ok {
			res = append(res, hash[v])
		} else {
			res = append(res, -1)
		}
	}

	return res
}
