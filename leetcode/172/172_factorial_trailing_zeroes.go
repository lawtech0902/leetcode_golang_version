/*
__author__ = 'lawtech'
__date__ = '2020/03/17 9:47 上午'
*/

package _172

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	} else {
		return n/5 + trailingZeroes(n/5)
	}
}
