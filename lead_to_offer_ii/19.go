package lead_to_offer_ii

// 双指针
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return strictPalindrome(s[l+1:r+1]) || strictPalindrome(s[l:r])
		}

		l++
		r--
	}

	return true
}

func strictPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
