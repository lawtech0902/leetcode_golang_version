/*
__author__ = 'lawtech'
__date__ = '2020/05/13 9:39 上午'
*/

package _400

func findNthDigit(n int) int {
	start, size, step := 1, 1, 9
	for n > size*step {
		n, size, step, start = n-(size*step), size+1, step*10, start*10
	}

	index := (n - 1) % size
	count := (n - 1) / size
	num := start + count
	for size-index > 1 {
		num /= 10
		index++
	}

	return num % 10
}
