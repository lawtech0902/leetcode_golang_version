/*
__author__ = 'lawtech'
__date__ = '2020/08/11 10:49 上午'
*/

package _64

func sumNums(n int) int {
	if n == 1 {
		return n
	}

	return n + sumNums(n-1)
}
