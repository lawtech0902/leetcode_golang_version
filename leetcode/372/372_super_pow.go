/*
__author__ = 'lawtech'
__date__ = '2020/05/07 1:49 下午'
*/

package _372

const base = 1337

func superPow(a int, b []int) int {
	n := len(b)
	if n == 0 {
		return 1
	}

	lastDigit := b[n-1]
	return powMod(superPow(a, b[:n-1]), 10) * powMod(a, lastDigit) % base
}

func powMod(a, k int) int {
	a %= base
	res := 1
	for i := 0; i < k; i++ {
		res = res * a % base
	}

	return res
}
