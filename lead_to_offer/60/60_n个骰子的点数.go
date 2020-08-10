/*
__author__ = 'lawtech'
__date__ = '2020/08/10 11:16 上午'
*/

package _60

import "math"

// dp 空间未优化
func twoSum(n int) []float64 {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 6*n+1)
	}

	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i + 1; j <= 6*i; j++ {
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}

	var res []float64
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[n][i])/math.Pow(6, float64(n)))
	}

	return res
}

// dp 空间优化
func twoSum1(n int) []float64 {
	dp := make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			dp[j] = 0
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					dp[j] += dp[j-k]
				}
			}
		}
	}

	var res []float64
	for k := n; k <= 6*n; k++ {
		res = append(res, float64(dp[k])/math.Pow(6, float64(n)))
	}

	return res
}
