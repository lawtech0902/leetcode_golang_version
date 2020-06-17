/*
__author__ = 'lawtech'
__date__ = '2020/06/17 10:35 上午'
*/

package _561

import "sort"

// 排序
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}
