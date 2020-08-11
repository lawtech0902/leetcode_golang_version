/*
__author__ = 'lawtech'
__date__ = '2020/08/11 10:54 上午'
*/

package _65

func add(a, b int) int {
	for b != 0 { // 进位和为0跳出循环
		c := (a & b) << 1 // 进位
		a ^= b            // 非进位和
		b = c
	}

	return a
}
