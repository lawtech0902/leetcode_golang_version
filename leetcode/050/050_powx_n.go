/*
__author__ = 'lawtech'
__date__ = '2020/1/15 4:06 下午'
*/

package _50

func myPow(x float64, n int) float64 {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	case -1:
		return 1 / x
	}

	b := myPow(x, n/2)
	l := myPow(x, n%2)
	return b * b * l
}
