/*
__author__ = 'lawtech'
__date__ = '2020/07/15 2:15 下午'
*/

package _712

// dp
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		dp[i][n] = dp[i+1][n] + int(s1[i])
	}

	for j := n - 1; j >= 0; j-- {
		dp[m][j] = dp[m][j+1] + int(s2[j])
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 转换为求最大公共子串，并要求这个子串的ascii值最大
func minimumDeleteSum1(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	totalSum := 0
	for i := 0; i < m; i++ {
		totalSum += int(s1[i])
	}

	for j := 0; j < n; j++ {
		totalSum += int(s2[j])
	}

	if m == 0 || n == 0 {
		return totalSum
	}

	commonSum := make([][]int, m+1)
	for i := range commonSum {
		commonSum[i] = make([]int, n+1)
	}

	maxCommonSum := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				commonSum[i][j] = commonSum[i-1][j-1] + int(s1[i-1])
			} else {
				commonSum[i][j] = max(commonSum[i-1][j], commonSum[i][j-1])
			}

			maxCommonSum = max(maxCommonSum, commonSum[i][j])
		}
	}

	return totalSum - 2*maxCommonSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
