/*
__author__ = 'lawtech'
__date__ = '2020/04/01 10:37 上午'
*/

package _221

import "math"

// dp
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	var maxSize int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}

			if maxSize < dp[i][j] {
				maxSize = dp[i][j]
			}
		}
	}

	return maxSize * maxSize
}

func min(nums ...int) int {
	min := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}
