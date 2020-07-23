/*
__author__ = 'lawtech'
__date__ = '2020/07/23 1:30 下午'
*/

package _775

import "math"

func isIdealPermutation(A []int) bool {
	for i := 0; i < len(A); i++ {
		if int(math.Abs(float64(A[i]-i))) > 1 {
			return false
		}
	}

	return true
}
