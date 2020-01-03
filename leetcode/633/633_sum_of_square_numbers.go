package _633

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/24 下午6:03'
*/

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))

	for i <= j {
		powSum := i*i + j*j
		if powSum == c {
			return true
		} else if powSum > c {
			j--
		} else {
			i++
		}
	}

	return false
}
