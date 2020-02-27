/*
__author__ = 'lawtech'
__date__ = '2020/02/26 5:16 下午'
*/

package _115

import "math"

func numDistinct(s string, t string) int {
	// dp[i][j]表示S[0...i-1]中有多少子串是T[0...j-1]。
	sSize, tSize := len(s), len(t)

	dp := make([][]int, sSize+1)
	for i := range dp {
		dp[i] = make([]int, tSize+1)

		for j := range dp[i] {
			dp[i][j] = 0
		}

		dp[i][0] = 1
	}

	for i := 1; i < sSize+1; i++ {
		for j := 1; j < int(math.Min(float64(i+1), float64(tSize+1))); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[sSize][tSize]
}
