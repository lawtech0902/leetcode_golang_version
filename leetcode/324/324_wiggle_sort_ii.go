/*
__author__ = 'lawtech'
__date__ = '2020/04/23 11:17 上午'
*/

package _324

import "sort"

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	sort.Ints(nums)
	mid := len(nums)/2 + len(nums)%2
	var res []int

	for i := 0; i < mid; i++ {
		res = append(res, nums[mid-i-1])
		if mid+i < len(nums) {
			res = append(res, nums[len(nums)-i-1])
		}
	}

	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
}
