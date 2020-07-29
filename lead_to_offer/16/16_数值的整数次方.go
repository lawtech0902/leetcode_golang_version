/*
__author__ = 'lawtech'
__date__ = '2020/07/29 10:26 上午'
*/

package _16

// 快速幂
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	b := n
	res := 1.0
	if b < 0 {
		x = 1 / x
		b = -b
	}

	for b > 0 {
		if b&1 == 1 {
			res *= x
		}

		x *= x
		b >>= 1
	}

	return res
}
