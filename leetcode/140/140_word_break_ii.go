/*
__author__ = 'lawtech'
__date__ = '2020/03/09 1:36 下午'
*/

package _140

func wordBreak(s string, wordDict []string) []string {
	// dp + dfs剪枝
	var (
		res []string
		dfs func(s string, wordDict []string, subStr string)
	)

	dfs = func(s string, wordDict []string, subStr string) {
		if check(s, wordDict) {
			if len(s) == 0 {
				res = append(res, subStr[1:])
			}

			for i := 1; i < len(s)+1; i++ {
				if inWordDict(wordDict, s[:i]) {
					dfs(s[i:], wordDict, subStr+" "+s[:i])
				}
			}
		}
	}

	dfs(s, wordDict, "")
	return res
}

func check(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] == true && inWordDict(wordDict, s[j:i+1]) {
				dp[i+1] = true
			}
		}
	}

	return dp[n]
}

func inWordDict(wordDict []string, s string) bool {
	for _, word := range wordDict {
		if word == s {
			return true
		}
	}

	return false
}
