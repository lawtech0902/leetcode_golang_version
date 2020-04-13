/*
__author__ = 'lawtech'
__date__ = '2020/04/13 4:36 下午'
*/

package _258

func addDigits(num int) int {
	if num <= 9 {
		return num
	}

	return (num-10)%9 + 1
}
