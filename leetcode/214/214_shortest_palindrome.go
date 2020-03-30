/*
__author__ = 'lawtech'
__date__ = '2020/03/27 4:59 下午'
*/

package _214

// 双指针法
func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	index := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[index] == s[i] {
			index++
		}
	}

	if index == len(s) {
		return s
	}

	return reverse(s[index:]) + shortestPalindrome(s[:index]) + s[index:]
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
