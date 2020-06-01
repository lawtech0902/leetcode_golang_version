/*
__author__ = 'lawtech'
__date__ = '2020/06/01 9:47 上午'
*/

package _486

// 递归
func PredictTheWinner(nums []int) bool {
	var winner func(start, end, turn int) int
	winner = func(start, end, turn int) int {
		if start == end {
			return turn * nums[start]
		}

		a := turn*nums[start] + winner(start+1, end, -turn)
		b := turn*nums[end] + winner(start, end-1, -turn)
		return turn * max(turn*a, turn*b)
	}

	return winner(0, len(nums)-1, 1) >= 0
}

// dp
func PredictTheWinner1(nums []int) bool {
	dp := make([]int, len(nums))
	for start := len(nums); start >= 0; start-- {
		for end := start + 1; end < len(nums); end++ {
			a := nums[start] - dp[end]
			b := nums[end] - dp[end-1]
			dp[end] = max(a, b)
		}
	}

	return dp[len(nums)-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
