/*
__author__ = 'lawtech'
__date__ = '2020/1/9 5:23 下午'
*/

package _16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.Inf(1)
	size := len(nums)
	res := 0

	for i, num := range nums {
		left, right := i+1, size-1
		for left < right {
			total := num + nums[left] + nums[right]
			diff := math.Abs(float64(total - target))
			if diff < minDiff {
				minDiff = diff
				res = total
			}

			if total == target {
				return total
			} else if total < target {
				left += 1
			} else {
				right -= 1
			}
		}
	}

	return res
}
