/*
__author__ = 'lawtech'
__date__ = '2020/03/27 10:51 上午'
*/

package _213

import "math"

// 划分为198的两个单排问题，再求最优解
func rob(nums []int) int {
	if len(nums) > 1 {
		return max(myRob(nums[:len(nums)-1]), myRob(nums[1:]))
	} else if len(nums) == 0 {
		return 0
	} else {
		return nums[0]
	}
}

func myRob(nums []int) int {
	pre, cur := 0, 0
	for _, num := range nums {
		pre, cur = cur, int(math.Max(float64(pre+num), float64(cur)))
	}

	return cur
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
