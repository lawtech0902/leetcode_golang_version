/*
__author__ = 'lawtech'
__date__ = '2020/1/17 4:44 下午'
*/

package _90

import (
	"reflect"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	// 效率太差
	sort.Ints(nums)
	var res [][]int
	dfs(nums, 0, []int{}, &res)
	return res
}

func dfs(nums []int, start int, valSlice []int, res *[][]int) {
	if !sliceInRes(valSlice, *res) {
		*res = append(*res, valSlice)
	}

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

func sliceInRes(slice []int, res [][]int) bool {
	for _, val := range res {
		if reflect.DeepEqual(val, slice) {
			return true
		}
	}

	return false
}
