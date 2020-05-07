/*
__author__ = 'lawtech'
__date__ = '2020/05/07 10:52 上午'
*/

package _368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	parent := make([]int, n)
	maxNum, maxIndex := 0, 0

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if nums[j]%nums[i] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				parent[i] = j
				if dp[i] > maxNum {
					maxNum = dp[i]
					maxIndex = i
				}
			}
		}
	}

	res := make([]int, 0, n)
	curIndex := maxIndex
	for i := 0; i < maxNum; i++ {
		res = append(res, nums[curIndex])
		curIndex = parent[curIndex]
	}

	return res
}
