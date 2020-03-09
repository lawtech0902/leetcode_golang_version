/*
__author__ = 'lawtech'
__date__ = '2020/03/09 10:54 上午'
*/

package _137

func singleNumber(nums []int) int {
	// 异或
	a, b := 0, 0
	for _, v := range nums {
		a = (a ^ v) & ^b
		b = (b ^ v) & ^a
	}
	return a
}
