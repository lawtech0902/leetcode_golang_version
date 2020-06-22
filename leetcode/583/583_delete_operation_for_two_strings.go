/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-01-03 16:19:42
 */
package _583

import "math"

// 转化为求两个字符串的最长公共子序列问题
func minDistance(word1 string, word2 string) int {
	return len(word1) + len(word2) - 2*lengthOfLCS(word1, word2)
}

func lengthOfLCS(word1, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return dp[n1][n2]
}
