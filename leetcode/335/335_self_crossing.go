/*
__author__ = 'lawtech'
__date__ = '2020/04/27 4:42 下午'
*/

package _335

func isSelfCrossing(x []int) bool {
	b, c, d, e, f := 0, 0, 0, 0, 0
	for _, a := range x {
		if d >= b && b > 0 && (a >= c || (a >= (c-e) && (c-e) >= 0) && f >= d-b) {
			return true
		}

		b, c, d, e, f = a, b, c, d, e
	}

	return false
}
