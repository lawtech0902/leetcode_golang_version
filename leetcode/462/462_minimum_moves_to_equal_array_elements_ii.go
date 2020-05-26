/*
__author__ = 'lawtech'
__date__ = '2020/05/26 9:55 上午'
*/

package _462

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	res := 0
	for i < j {
		res += nums[j] - nums[i]
		i++
		j--
	}

	return res
}
