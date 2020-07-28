/*
__author__ = 'lawtech'
__date__ = '2020/07/27 3:30 下午'
*/

package _10

func fib(n int) int {
	a, b, sum := 0, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a, b = b, sum
	}

	return a
}
