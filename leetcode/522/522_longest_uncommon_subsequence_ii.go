/*
__author__ = 'lawtech'
__date__ = '2020/06/05 1:42 下午'
*/

package _522

import "math"

func findLUSlength(strs []string) int {
	res := -1
	for i, j := 0, 0; i < len(strs); i++ {
		for j = 0; j < len(strs); j++ {
			if j == i {
				continue
			}

			if isSubsequence(strs[i], strs[j]) {
				break
			}
		}

		if j == len(strs) {
			res = int(math.Max(float64(res), float64(len(strs[i]))))
		}
	}

	return res
}

func isSubsequence(s1, s2 string) bool {
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s1[j] == s2[i] {
			j++
		}
	}

	return j == len(s1)
}
