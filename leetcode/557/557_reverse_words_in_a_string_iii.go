/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-16 13:52:34
 */

package _557

import "strings"

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i := 0; i < len(ss); i++ {
		ss[i] = reverse(ss[i])
	}

	return strings.Join(ss, " ")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
