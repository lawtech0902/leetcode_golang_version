/*
__author__ = 'lawtech'
__date__ = '2020/03/23 10:14 上午'
*/

package _198

import "math"

// dp优化
func rob(nums []int) int {
	prevMax, curMax := 0, 0
	for _, num := range nums {
		temp := curMax
		curMax = int(math.Max(float64(prevMax+num), float64(curMax)))
		prevMax = temp
	}

	return curMax
}

// 原始dp
func rob1(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	dp := make([]int, size+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= size; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i-1])))
	}

	return dp[size]
}
