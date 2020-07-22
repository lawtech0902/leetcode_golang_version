/*
__author__ = 'lawtech'
__date__ = '2020/07/22 2:42 下午'
*/

package _762

import "math/bits"

func countPrimeSetBits(L, R int) int {
	res := 0
	for i := L; i <= R; i++ {
		if isSmallPrime(bits.OnesCount(uint(i))) {
			res++
		}
	}

	return res
}

func isSmallPrime(x int) bool {
	return x == 2 || x == 3 || x == 5 || x == 7 ||
		x == 11 || x == 13 || x == 17 || x == 19
}
