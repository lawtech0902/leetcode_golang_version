/*
__author__ = 'lawtech'
__date__ = '2020/05/06 4:58 下午'
*/

package _357

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	res, muls := 10, 9
	for i := 1; i < min(n, 10); i++ {
		muls *= 10 - i
		res += muls
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
