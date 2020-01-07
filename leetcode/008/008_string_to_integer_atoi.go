/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-01-03 16:19:42
 */
package _08

import (
	"math"
	"strings"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 上午12:05'
*/

func myAtoi(str string) int {
	// 跳过首部空格
	// for idx < size && (str[idx] == ' ' || str[idx] == '\t') {
	// 	idx++
	// }
	str = strings.TrimLeft(str, " \t")

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
