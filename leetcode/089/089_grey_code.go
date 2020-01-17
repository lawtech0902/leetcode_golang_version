/*
__author__ = 'lawtech'
__date__ = '2020/1/17 4:39 下午'
*/

package _89

import "math"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	size := int(math.Exp2(float64(n)))
	data := make([]int, size)
	data[0] = 0
	base := 1

	for digit := 0; digit < n; digit++ { // 控制位数
		for i := 0; i < base; i++ { // 每增加一位,前面所有的数高位+1 反向
			data[base+i] = data[base-i-1] + base
		}
		base *= 2
	}
	return data
}
