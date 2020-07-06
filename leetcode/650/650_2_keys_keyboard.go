/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-01-03 16:19:42
 */

package _650

// dp
func minSteps(n int) int {
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := i - 1; j > 1; j-- {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 素数分解
func minSteps1(n int) int {
	res, d := 0, 2
	for n > 1 {
		for n%d == 0 {
			res += d
			n /= d
		}

		d++
	}

	return res
}
