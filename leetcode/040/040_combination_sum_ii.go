/*
__author__ = 'lawtech'
__date__ = '2020/1/13 4:09 下午'
*/

package _40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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

	dfs(candidates[1:], target-candidates[0], append(valSlice, candidates[0]), res)

	dfs(next(candidates), target, valSlice, res)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}

	return candidates[i+1:]
}
