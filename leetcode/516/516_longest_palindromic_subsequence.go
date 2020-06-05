/*
__author__ = 'lawtech'
__date__ = '2020/06/05 9:56 上午'
*/

package _516

// dp[i][j] 表示 s 的第 i 个字符到第 j 个字符组成的子串中，最长的回文序列长度是多少。
func longestPalindromeSubseq(s string) int {
	size := len(s)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	for i := size - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
