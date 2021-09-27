package lead_to_offer_ii

import "unicode"

// 双指针
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !unicode.IsDigit(rune(s[l])) && !unicode.IsLetter(rune(s[l])) {
			l++
			continue
		}

		if !unicode.IsDigit(rune(s[r])) && !unicode.IsLetter(rune(s[r])) {
			r--
			continue
		}

		if unicode.IsDigit(rune(s[l])) {
			if s[l] != s[r] {
				return false
			}
		} else {
			if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
				return false
			}
		}

		l++
		r--
	}

	return true
}
