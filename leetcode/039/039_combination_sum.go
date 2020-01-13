/*
__author__ = 'lawtech'
__date__ = '2020/1/13 2:39 下午'
*/

package _39

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var (
		res      [][]int
		valSlice []int
	)

	sort.Ints(candidates)
	dfs(candidates, target, valSlice, &res)
	return res
}

func dfs(candidates []int, target int, valSlice []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, valSlice)
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	valSlice = valSlice[:len(valSlice):len(valSlice)]

	dfs(candidates, target-candidates[0], append(valSlice, candidates[0]), res)

	dfs(candidates[1:], target, valSlice, res)
}
