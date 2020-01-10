/*
__author__ = 'lawtech'
__date__ = '2020/1/10 2:34 下午'
*/

package _29

import "math"

func divide(dividend int, divisor int) int {
	// 位运算
	var flag bool
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		flag = false
	} else {
		flag = true
	}

	divd, divs := int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	res := 0

	for divd >= divs {
		tmp, i := divs, 1
		for divd-tmp >= 0 {
			divd -= tmp
			res += i
			i <<= 1
			tmp <<= 1
		}
	}

	if flag == false {
		res = -1 * res
	}

	return int(math.Min(math.MaxInt32, math.Max(float64(res), math.MinInt32)))
}
