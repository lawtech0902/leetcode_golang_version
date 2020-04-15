/*
__author__ = 'lawtech'
__date__ = '2020/04/15 1:46 下午'
*/

package _279

import "math"

// dp
func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-j*j]+1)))
		}
	}

	return dp[n]
}

func isSquare(n int) bool {
	sqrtN := int(math.Sqrt(float64(n)))
	return sqrtN*sqrtN == n
}

// 四平方和定理
func numSquares1(n int) int {
	if n == 0 {
		return 0
	}

	if isSquare(n) {
		return 1
	}

	for (n & 3) == 0 {
		n >>= 2
	}

	if (n & 7) == 7 {
		return 4
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3
}
