/*
__author__ = 'lawtech'
__date__ = '2020/05/12 4:05 下午'
*/

package _397

// 迭代
func integerReplacement(n int) int {
	count := 0
	for n > 1 {
		if n == 3 {
			n -= 1
		} else if n%4 == 1 {
			n -= 1
		} else if n%4 == 3 {
			n += 1
		} else {
			n /= 2
		}

		count++
	}

	return count
}
