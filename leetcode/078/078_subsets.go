/*
__author__ = 'lawtech'
__date__ = '2020/1/17 11:23 上午'
*/

package _78

import "sort"

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	dfs(nums, 0, []int{}, &res)
	return res
}

func dfs(nums []int, start int, valSlice []int, res *[][]int) {
	*res = append(*res, valSlice)
	if start == len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		tmp := make([]int, len(valSlice)+1)
		copy(tmp, valSlice)
		tmp[len(valSlice)] = nums[i]
		dfs(nums, i+1, tmp, res)
	}
}
