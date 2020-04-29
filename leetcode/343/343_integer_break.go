/*
__author__ = 'lawtech'
__date__ = '2020/04/29 10:48 上午'
*/

package _343

// dp
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(j, dp[j])*max(i-j, dp[i-j]))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// math
func integerBreak1(n int) int {
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	product := 1
	for n > 4 {
		product *= 3
		n -= 3
	}

	return product * n
}
