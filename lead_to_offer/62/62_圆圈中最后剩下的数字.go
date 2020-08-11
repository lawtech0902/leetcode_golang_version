/*
__author__ = 'lawtech'
__date__ = '2020/08/11 9:41 上午'
*/

package _62

func lastRemaining(n, m int) int {
	if n == 1 {
		return 0
	}

	return (m + lastRemaining(n-1, m)) % n
}
