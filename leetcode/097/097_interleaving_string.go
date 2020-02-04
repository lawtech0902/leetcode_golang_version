/*
__author__ = 'lawtech'
__date__ = '2020/2/4 2:48 下午'
*/

package _97

func isInterleave(s1 string, s2 string, s3 string) bool {
	// dp[i][j]表示s1[0...i-1]和s2[0...j-1]是否可以拼接为s3[0...i+j-1]，可以拼接为true，不可以拼接为false
	// 遇到字符串问题先想dp，不行再dfs
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true

	for i := 1; i < len(s1)+1; i++ {
		dp[i][0] = dp[i-1][0] && s3[i-1] == s1[i-1]
	}

	for i := 1; i < len(s2)+1; i++ {
		dp[0][i] = dp[0][i-1] && s3[i-1] == s2[i-1]
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[len(s1)][len(s2)]
}
