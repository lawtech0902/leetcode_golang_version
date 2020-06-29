/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-29 14:45:11
 */

package _611

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0

	for i := len(nums) - 1; i > 1; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				res += (r - 1) - l + 1
				r -= 1
			} else {
				l += 1
			}
		}
	}

	return res
}
