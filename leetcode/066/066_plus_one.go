/*
__author__ = 'lawtech'
__date__ = '2020/1/16 3:32 下午'
*/

package _66

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits = append(digits, 0)
	copy(digits[1:], digits[:])
	digits[0] = 1

	return digits
}
