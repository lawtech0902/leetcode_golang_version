/*
__author__ = 'lawtech'
__date__ = '2020/05/18 5:27 下午'
*/

package _434

import "strings"

// built-in
func countSegments(s string) int {
	return len(strings.Fields(s))
}

func countSegments1(s string) int {
	var count int
	for i := range s {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			count++
		}
	}

	return count
}
