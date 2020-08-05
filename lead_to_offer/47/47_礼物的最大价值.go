/*
__author__ = 'lawtech'
__date__ = '2020/08/05 10:47 上午'
*/

package _47

import "math"

func maxValue(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	// dp[i][j]表示从grid[0][0]到grid[i - 1][j - 1]时的最大价值
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}

	return dp[row][col]
}

func max(values ...int) int {
	maxValue := math.MinInt32
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

// 由于 dp[i][j] 只与 dp[i-1][j] , dp[i][j-1] , grid[i][j] 有关系，
// 因此可以将原矩阵 grid 用作 dp 矩阵，即直接在 grid 上修改即可。
func maxValue1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 初始化第一行
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// 初始化第一列
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += max(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[m-1][n-1]
}
