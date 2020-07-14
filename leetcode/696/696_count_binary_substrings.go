/*
__author__ = 'lawtech'
__date__ = '2020/07/14 11:29 上午'
*/

package _696

func countBinarySubstrings(s string) int {
	cur := 1
	pre := 0
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			cur++
		} else {
			res += min(cur, pre)
			pre = cur
			cur = 1
		}
	}

	return res + min(cur, pre)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
