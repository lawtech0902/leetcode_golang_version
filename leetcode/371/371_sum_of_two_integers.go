/*
__author__ = 'lawtech'
__date__ = '2020/05/07 1:35 下午'
*/

package _371

func getSum(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	var lower int
	for {
		lower = a ^ b
		carrier := a & b
		if carrier == 0 {
			break
		}

		a = lower
		b = carrier << 1
	}

	return lower
}
