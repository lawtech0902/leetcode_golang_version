/*
__author__ = 'lawtech'
__date__ = '2020/1/17 3:56 下午'
*/

package _87

import (
	"sort"
	"strings"
)

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	if sortString(s1) != sortString(s2) {
		return false
	}

	size := len(s1)
	for i := 1; i < size; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[size-i:]) && isScramble(s1[i:], s2[:size-i]) {
			return true
		}
	}

	return false
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
