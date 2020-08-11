/*
__author__ = 'lawtech'
__date__ = '2020/08/11 4:44 下午'
*/

package _67

import (
	"math"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimLeft(str, "\t")
	res, sign, size, idx := 0, 1, len(str), 0
	// +/-
	if idx < size {
		if str[idx] == '+' {
			sign = 1
			idx++
		} else if str[idx] == '-' {
			sign = -1
			idx++
		}
	}

	// 转换
	for idx < size && str[idx] >= '0' && str[idx] <= '9' {
		res = res*10 + int(str[idx]) - '0'
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*res < math.MinInt32 {
			return math.MinInt32
		}

		idx++
	}

	return res * sign
}
