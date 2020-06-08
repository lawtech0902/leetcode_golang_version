/*
__author__ = 'lawtech'
__date__ = '2020/06/08 10:35 上午'
*/

package _525

import "math"

// hashmap
func findMaxLength(nums []int) int {
	hash := make(map[int]int)
	hash[0] = -1
	maxLen, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}

		if val, ok := hash[count]; ok {
			maxLen = int(math.Max(float64(maxLen), float64(i-val)))
		} else {
			hash[count] = i
		}
	}

	return maxLen
}

// 暴力
func findMaxLength1(nums []int) int {
	maxLen := 0
	for start := 0; start < len(nums); start++ {
		zeroes, ones := 0, 0
		for end := start; end < len(nums); end++ {
			if nums[end] == 0 {
				zeroes++
			} else {
				ones++
			}

			if zeroes == ones {
				maxLen = int(math.Max(float64(maxLen), float64(end-start+1)))
			}
		}
	}

	return maxLen
}
