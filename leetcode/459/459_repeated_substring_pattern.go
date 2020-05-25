/*
__author__ = 'lawtech'
__date__ = '2020/05/25 10:11 上午'
*/

package _459

import "strings"

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	ss := (s + s)[1 : len(s)*2-1]
	return strings.Contains(ss, s)
}
