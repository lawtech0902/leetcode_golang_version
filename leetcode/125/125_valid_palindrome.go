/*
__author__ = 'lawtech'
__date__ = '2020/03/03 3:44 下午'
*/

package _125

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		if !('a' <= s[l] && s[l] <= 'z' || '0' <= s[l] && s[l] <= '9') {
			l++
			continue
		}

		if !('a' <= s[r] && s[r] <= 'z' || '0' <= s[r] && s[r] <= '9') {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
