package 剑指offer

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/26 下午9:19'
*/

func jumpFloorII(n int) int {
	return int(math.Pow(float64(2), float64(n-1)))
}
