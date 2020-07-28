/*
__author__ = 'lawtech'
__date__ = '2020/07/28 3:43 下午'
*/

package _14

import "math"

// 数学规则推导
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	a := n / 3
	b := n % 3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 4
	} else {
		return int(math.Pow(3, float64(a))) * 2
	}
}
