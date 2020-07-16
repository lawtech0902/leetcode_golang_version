/*
__author__ = 'lawtech'
__date__ = '2020/07/16 10:42 上午'
*/

package _717

func isOneBitCharacter(bits []int) bool {
	ones := 0
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 1 {
			ones++
		} else {
			break
		}
	}

	return ones%2 == 0
}
