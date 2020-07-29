/*
__author__ = 'lawtech'
__date__ = '2020/07/29 2:05 下午'
*/

package _19

// dp
func isMatch(s, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 分成非空正则和空正则两种
			if j == 0 {
				dp[i][j] = i == 0
			} else {
				// 非空正则分为两种情况 * 和 非*
				if p[j-1] != '*' {
					if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
						dp[i][j] = dp[i-1][j-1]
					}
				} else {
					// 碰到*了分为看和不看两种情况
					// 不看
					if j >= 2 {
						dp[i][j] = dp[i][j] || dp[i][j-2]
					}

					// 看
					if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
						dp[i][j] = dp[i][j] || dp[i-1][j]
					}
				}
			}
		}
	}

	return dp[n][m]
}
