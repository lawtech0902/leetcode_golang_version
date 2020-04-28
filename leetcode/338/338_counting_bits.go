/*
__author__ = 'lawtech'
__date__ = '2020/04/28 10:40 上午'
*/

package _338

import "math"

func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	if num == 0 {
		return res
	}

	res[1] = 1
	power := 0
	for i := 2; i <= num; i++ {
		if i == (int)(math.Pow(2, (float64)(power+1))) {
			power++
			res[i] = 1
		} else {
			res[i] = 1 + res[i-(int)(math.Pow(2, (float64)(power)))]
		}
	}

	return res
}
