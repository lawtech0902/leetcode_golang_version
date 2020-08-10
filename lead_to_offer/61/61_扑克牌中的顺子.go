/*
__author__ = 'lawtech'
__date__ = '2020/08/10 3:04 下午'
*/

package _61

import (
	"math"
	"sort"
)

// 哈希 + 遍历
func isStraight(nums []int) bool {
	repeat := make(map[int]bool, 14)
	ma, mi := 0, 14
	for _, num := range nums {
		if num == 0 {
			continue
		}

		ma = max(ma, num)
		mi = min(mi, num)
		if repeat[num] {
			return false
		}

		repeat[num] = true
	}

	return ma-mi < 5
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

func min(values ...int) int {
	minValue := math.MaxInt32
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}

// 排序 + 遍历
func isStraight1(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}

	return nums[4]-nums[joker] < 5
}
