/*
__author__ = 'lawtech'
__date__ = '2018/8/18 上午12:57'
*/

package _15

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)
	size := len(nums)

	if size <= 2 {
		return res
	}

	for i := range nums[:size-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := nums[i] * -1
		left, right := i+1, size-1
		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				right -= 1
				left += 1

				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}
			} else if nums[left]+nums[right] > target {
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return res
}
