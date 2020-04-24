/*
__author__ = 'lawtech'
__date__ = '2020/04/24 10:52 上午'
*/

package _326

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	return 4052555153018976267%n == 0
}
