/*
__author__ = 'lawtech'
__date__ = '2020/03/30 1:26 下午'
*/

package _216

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	dfs(k, n, 1, &res, nil)
	return res
}

func dfs(k, n, start int, res *[][]int, cur []int) {
	if k == 0 && n == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	if k == 0 || n < 0 {
		return
	}

	for i := start; i <= 9; i++ {
		dfs(k-1, n-i, i+1, res, append(cur, i))
	}
}
