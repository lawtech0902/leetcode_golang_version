/*
__author__ = 'lawtech'
__date__ = '2020/05/08 11:09 上午'
*/

package _375

import "math"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	return helper(1, n, &dp)
}

func helper(l, r int, dp *[][]int) int {
	if l >= r {
		return 0
	}

	if (*dp)[l][r] != math.MaxInt32 {
		return (*dp)[l][r]
	}

	for i := l; i <= r; i++ {
		tmp := max(i+helper(l, i-1, dp), i+helper(i+1, r, dp))
		(*dp)[l][r] = min((*dp)[l][r], tmp)
	}

	return (*dp)[l][r]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
