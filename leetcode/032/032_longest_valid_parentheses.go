/*
__author__ = 'lawtech'
__date__ = '2020/1/10 3:38 下午'
*/

package _32

import "math"

func longestValidParentheses(s string) int {
	maxLen := 0
	size := len(s)
	dp := make([]int, size)

	for i := 1; i < size; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if (i - dp[i-1]) >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}

			maxLen = int(math.Max(float64(maxLen), float64(dp[i])))
		}
	}

	return maxLen
}
