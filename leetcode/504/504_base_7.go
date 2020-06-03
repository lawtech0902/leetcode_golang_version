/*
__author__ = 'lawtech'
__date__ = '2020/06/03 1:59 下午'
*/

package _504

import (
	"math"
	"strconv"
)

// recursive
func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(-1*num)
	}

	if num < 7 {
		return strconv.Itoa(num)
	}

	return convertToBase7(num/7) + strconv.Itoa(num%7)
}

// iterative
func convertToBase71(num int) string {
	if num == 0 {
		return "0"
	}

	var res string
	n := int(math.Abs(float64(num)))
	for n != 0 {
		res = strconv.Itoa(n%7) + res
		n /= 7
	}

	if num < 0 {
		return "-" + res
	}

	return res
}
