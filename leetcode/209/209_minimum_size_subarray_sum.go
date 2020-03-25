/*
__author__ = 'lawtech'
__date__ = '2020/03/25 11:27 上午'
*/

package _209

import "math"

// 双指针 滑动窗口
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	res := math.MaxInt32
	left := 0
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= s {
			res = int(math.Min(float64(res), float64(i+1-left)))
			sum -= nums[left]
			left++
		}
	}

	if res != math.MaxInt32 {
		return res
	}

	return 0
}
