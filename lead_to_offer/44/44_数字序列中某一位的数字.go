/*
__author__ = 'lawtech'
__date__ = '2020/08/04 2:04 下午'
*/

package _44

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
