/*
__author__ = 'lawtech'
__date__ = '2020/07/09 2:44 下午'
*/

package _680

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
