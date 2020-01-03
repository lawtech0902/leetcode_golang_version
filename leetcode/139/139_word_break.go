package _139

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 上午12:47'
*/

/*
dp[i]表示字符串s[:i]能否拆分成符合要求的子字符串。
我们可以看出，如果s[j:i]在给定的字符串组中，
且dp[j]为True（即字符串s[:j]能够拆分成符合要求的子字符串），
那么此时dp[i]也就为True了。按照这种递推关系，我们就可以判断目标字符串能否成功拆分。
*/

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] == true && inWordDict(wordDict, s[j:i+1]) == true {
				dp[i+1] = true
			}
		}
	}

	return dp[n]
}

func inWordDict(wordDict []string, s string) bool {
	for _, word := range wordDict {
		if s == word {
			return true
		}
	}

	return false
}
