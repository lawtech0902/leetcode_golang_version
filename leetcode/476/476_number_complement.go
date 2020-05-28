/*
__author__ = 'lawtech'
__date__ = '2020/05/28 10:14 上午'
*/

package _476

// e.g: 101（5)*111=010
func findComplement(num int) int {
	tmp := num
	num2 := 1
	for tmp > 0 {
		num2 <<= 1
		tmp >>= 1
	}

	num2 -= 1
	return num ^ num2
}
