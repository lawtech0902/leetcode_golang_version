/*
__author__ = 'lawtech'
__date__ = '2020/07/15 3:09 下午'
*/

package _713

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	cur, left := 1, 0
	res := 0
	for right := 0; right < len(nums); right++ {
		cur *= nums[right]
		for cur >= k {
			cur /= nums[left]
			left++
		}

		res += right - left + 1
	}

	return res
}
