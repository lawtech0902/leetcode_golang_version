/*
__author__ = 'lawtech'
__date__ = '2020/1/17 11:06 上午'
*/

package _77

func combine(n int, k int) [][]int {
	var res [][]int

	dfs(1, k, n, []int{}, &res)
	return res
}

func dfs(start, k, n int, valList []int, res *[][]int) {
	if len(valList) == k {
		*res = append(*res, valList)
		return
	}

	// 剪枝
	for i := start; i < n-(k-len(valList))+2; i++ {
		// slice 深拷贝 切记
		tmp := make([]int, len(valList)+1)
		copy(tmp, valList)
		tmp[len(valList)] = i

		dfs(i+1, k, n, tmp, res)
	}
}
