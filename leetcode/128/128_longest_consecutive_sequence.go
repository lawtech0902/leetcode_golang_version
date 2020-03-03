/*
__author__ = 'lawtech'
__date__ = '2020/03/03 4:47 下午'
*/

package _128

import (
	"math"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	longestStreak, curStreak := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				curStreak++
			} else {
				longestStreak = int(math.Max(float64(longestStreak), float64(curStreak)))
				curStreak = 1
			}
		}
	}

	return int(math.Max(float64(longestStreak), float64(curStreak)))
}
