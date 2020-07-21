/*
__author__ = 'lawtech'
__date__ = '2020/07/20 2:50 下午'
*/

package _740

import "math"

// dp 打家劫舍
func deleteAndEarn(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	maxNum := max(nums...)
	all := make([]int, maxNum+1)
	for _, num := range nums {
		all[num]++
	}

	dp := make([]int, maxNum+1)
	dp[1] = all[1]
	dp[2] = max(dp[1], all[2]*2)
	for i := 2; i <= maxNum; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+i*all[i])
	}

	return dp[maxNum]
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
