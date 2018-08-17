package _07

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/17 下午11:57'
*/

func reverse(x int) int {
	sign := 1

	// 处理负数
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		// 取出末尾
		temp := x % 10
		// 放入res头部
		res = temp + res*10
		// 去除x的末尾
		x = x / 10
	}

	res = sign * res

	// 处理溢出
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}
