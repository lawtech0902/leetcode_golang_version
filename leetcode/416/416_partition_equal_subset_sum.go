/*
__author__ = 'lawtech'
__date__ = '2018/8/21 上午12:09'
*/

package _416

import "sort"

// 典型的背包问题，在n个物品中选出一定的物品填满sum/2的背包
func canPartition(nums []int) bool {
	sort.Ints(nums)
	sums := 0
	for _, num := range nums {
		sums += num
	}

	if sums&1 != 0 {
		return false
	}

	target := sums / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for j := target; j > num-1; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
