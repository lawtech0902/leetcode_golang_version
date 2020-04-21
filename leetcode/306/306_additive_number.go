/*
__author__ = 'lawtech'
__date__ = '2020/04/21 11:16 上午'
*/

package _306

import "strconv"

func isAdditiveNumber(num string) bool {
	return helper(num, 0, 1, 0, 0)
}

func helper(num string, i, n, prev2, prev1 int) bool {
	if i == len(num) {
		return n > 3
	}

	for size := 1; size < len(num)-i+1; size++ {
		if size > 1 && num[i] == '0' {
			return false
		}

		v, _ := strconv.Atoi(num[i : i+size])
		if n > 2 && v != prev2+prev1 {
			continue
		}

		if helper(num, i+size, n+1, prev1, v) {
			return true
		}
	}

	return false
}
