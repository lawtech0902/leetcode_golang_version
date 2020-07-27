/*
__author__ = 'lawtech'
__date__ = '2020/07/27 10:42 上午'
*/

package _3

import "sort"

// 原地置换 O(n) O(1)
func findRepeatNumber(nums []int) int {
	for id, num := range nums {
		if id == num {
			continue
		}

		if nums[num] == num {
			return num
		}

		nums[num], nums[id] = nums[id], nums[num]
	}

	return 0
}

// 排序+遍历 O(nlogn) O(1)
func findRepeatNumber1(nums []int) int {
	sort.Ints(nums)
	for id := range nums {
		if id > 0 {
			if nums[id] == nums[id-1] {
				return nums[id]
			}
		}
	}

	return 0
}

// hash O(n) O(n)
func findRepeatNumber2(nums []int) int {
	hash := make(map[int]bool)
	for _, num := range nums {
		if hash[num] {
			return num
		}

		hash[num] = true
	}

	return 0
}
