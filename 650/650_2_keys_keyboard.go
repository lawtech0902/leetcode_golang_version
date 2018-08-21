package _650

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 下午11:45'
*/

func minSteps(n int) int {
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := i - 1; j > 1; j-- {
			if i%j == 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[j]+i/j)))
			}
		}
	}

	return dp[n]
}
