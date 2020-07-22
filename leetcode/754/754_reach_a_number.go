/*
__author__ = 'lawtech'
__date__ = '2020/07/22 10:20 上午'
*/

package _754

import "math"

func reachNumber(target int) int {
	target = int(math.Abs(float64(target)))
	k := 0
	for target > 0 {
		k++
		target -= k
	}

	if target%2 == 0 {
		return k
	}

	return k + 1 + k%2
}
