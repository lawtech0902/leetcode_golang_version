/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-02 14:37:07
 */

package _643

import "math"

// 滑动窗口
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	res := float64(sum)
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		res = math.Max(float64(res), float64(sum))
	}

	return res / float64(k)
}
