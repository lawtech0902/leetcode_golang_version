/*
__author__ = 'lawtech'
__date__ = '2020/07/10 4:17 下午'
*/

package _686

import "strings"

func repeatedStringMatch(A, B string) int {
	max := len(B)/len(A) + 2
	tmp := strings.Repeat(A, max)
	if !strings.Contains(tmp, B) {
		return -1
	}

	for {
		tmp = strings.Repeat(A, max)
		if !strings.Contains(tmp, B) {
			return max + 1
		}

		max--
	}
}
