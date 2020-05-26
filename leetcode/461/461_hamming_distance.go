/*
__author__ = 'lawtech'
__date__ = '2020/05/26 9:43 上午'
*/

package _461

func hammingDistance(x, y int) int {
	xor := x ^ y
	res := 0
	for xor != 0 {
		if xor&1 != 0 {
			res++
		}

		xor = xor >> 1
	}

	return res
}
