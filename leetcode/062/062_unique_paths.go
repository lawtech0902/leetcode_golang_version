/*
__author__ = 'lawtech'
__date__ = '2020/1/16 2:13 下午'
*/

package _62

// dp一维优化，看不懂
// func uniquePaths(m int, n int) int {
// 	cur := make([]int, n)
// 	for i := range cur {
// 		cur[i] = 1
// 	}
//
// 	for i := 1; i < m; i++ {
// 		for j := 1; j < n; j++ {
// 			cur[j] += cur[j-1]
// 		}
// 	}
//
// 	return cur[len(cur)-1]
// }

func uniquePaths(m int, n int) int {
	// dp: dp[i][j] = dp[i-1][j] + dp[i][j-1]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+1 < m {
				dp[i+1][j] += dp[i][j]
			}

			if j+1 < n {
				dp[i][j+1] += dp[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}
