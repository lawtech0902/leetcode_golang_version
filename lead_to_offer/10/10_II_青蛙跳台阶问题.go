/*
__author__ = 'lawtech'
__date__ = '2020/07/28 9:57 上午'
*/

package _10

func numWays(n int) int {
	a, b, sum := 1, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return a
}
