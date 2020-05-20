/*
__author__ = 'lawtech'
__date__ = '2020/05/20 3:27 下午'
*/

package _441

import "math"

func arrangeCoins(n int) int {
	return int((-1 + math.Sqrt(float64(1+8*n))) / 2)
}
