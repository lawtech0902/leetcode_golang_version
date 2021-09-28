package lead_to_offer_ii

// 双指针
func countSubstrings(s string) int {
	var (
		res   int
		count func(s string, start, end int)
	)

	count = func(s string, start, end int) {
		for start >= 0 && end < len(s) && s[start] == s[end] {
			res++
			start--
			end++
		}
	}

	for i := 0; i < len(s); i++ {
		count(s, i, i)
		count(s, i, i+1)
	}

	return res
}

// dp
func countSubstrings1(s string) int {
	res := 0
	dp := make([]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] == s[j] && (j-i <= 2 || dp[j-1]) {
				res++
				dp[j] = true
			} else {
				dp[j] = false
			}
		}
	}

	return res
}
