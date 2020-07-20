/*
__author__ = 'lawtech'
__date__ = '2020/07/20 1:31 下午'
*/

package _738

import "strconv"

// 贪心
func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}

	res := []byte(strconv.Itoa(N))
	size := len(res)
	for i := 0; i < size-1; {
		// 当前数比后一位大，当前就减一，可能造成前一位比当前位大，所以i后退
		if res[i] > res[i+1] {
			res[i]--
			for j := i + 1; j < size; j++ {
				res[j] = '9'
			}

			i--
			if i < 0 {
				break
			}
		} else {
			i++
		}
	}

	num, _ := strconv.Atoi(string(res))
	return num
}
