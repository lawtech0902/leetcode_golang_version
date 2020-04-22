/*
__author__ = 'lawtech'
__date__ = '2020/04/22 9:59 上午'
*/

package _313

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	idx := make([]int, len(primes)) // 质数因子的指针列表
	dp := make([]int, n)
	dp[0] = 1

	for k := 1; k < n; k++ {
		// 取出质数列表中的每个质数与各自指针对应的超级丑数相乘的最小值
		min := math.MaxInt32
		for i := 0; i < len(primes); i++ {
			tmp := dp[idx[i]] * primes[i]
			if min > tmp {
				min = tmp
			}
		}

		dp[k] = min

		// 若最小值等于该质数乘以dp[idx[i]]（第i个质数的指针所对应的超级丑数）则对应指针往后移动一步，i+1
		for i := 0; i < len(primes); i++ {
			if min == dp[idx[i]]*primes[i] {
				idx[i]++
			}
		}
	}

	return dp[n-1]
}
