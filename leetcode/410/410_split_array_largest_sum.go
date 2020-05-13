/*
__author__ = 'lawtech'
__date__ = '2020/05/13 3:20 下午'
*/

package _410

import "math"

// dp, 这个比较好理解
func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = 0
	sub := make([]int, n+1)
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sub[i]-sub[k]))
			}
		}
	}

	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 二分+贪心
func splitArray1(nums []int, m int) int {
	var max, sum int
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}

	if m == 1 {
		return sum
	}

	left, right := max, sum
	for left <= right {
		mid := (left + right) / 2
		if helper(mid, m, nums) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func helper(target, m int, nums []int) bool {
	var sum int
	count := 1
	for _, n := range nums {
		sum += n
		if sum > target {
			sum = n
			count++
			if count > m {
				return false
			}
		}
	}

	return true
}
