package _72

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 下午11:23'
*/

// dp[i][j]表示word1[0...i]到word2[0...j]的编辑距离
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 && n == 0 || word1 == word2 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])+1)
			}
		}
	}

	return dp[m][n]
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
