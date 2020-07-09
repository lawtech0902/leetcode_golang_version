/*
__author__ = 'lawtech'
__date__ = '2020/07/09 10:51 上午'
*/

package _674

func findLengthOfLCIS(nums []int) int {
	res, anchor := 0, 0
	for i := range nums {
		if i >= 1 && nums[i-1] >= nums[i] {
			anchor = i
		}

		res = max(res, i-anchor+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
