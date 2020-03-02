/*
__author__ = 'lawtech'
__date__ = '2020/03/02 11:22 上午'
*/

package _120

import "math"

func minimumTotal(triangle [][]int) int {
	/*
		自底向上
		DP方程：dp[m][n] = min(dp[m+1][n], dp[m+1][n+1])+triangle[m][n]
		dp[m][n]是指第m行第n个节点的最小路径和
	*/
	if triangle == nil {
		return 0
	}

	size := len(triangle)
	res := triangle[size-1]
	for i := size - 2; i > -1; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[j] = int(math.Min(float64(res[j]), float64(res[j+1]))) + triangle[i][j]
		}
	}

	return res[0]
}
