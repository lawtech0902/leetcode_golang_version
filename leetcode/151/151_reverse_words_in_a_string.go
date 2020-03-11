/*
__author__ = 'lawtech'
__date__ = '2020/03/11 11:42 上午'
*/

package _151

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	size := len(s)

	if size == 0 {
		return s
	}

	strs := strings.Fields(s)
	var res strings.Builder
	for i := len(strs) - 1; i >= 0; i-- {
		res.WriteString(strs[i] + " ")
	}

	return strings.TrimSpace(res.String())
}
