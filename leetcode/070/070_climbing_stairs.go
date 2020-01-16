/*
__author__ = 'lawtech'
__date__ = '2020/1/16 4:18 下午'
*/

package _70

func climbStairs(n int) int {
	// dp[i] = dp[i-1] + dp[i-2]
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = -1
	}
	dp[0], dp[1] = 1, 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
