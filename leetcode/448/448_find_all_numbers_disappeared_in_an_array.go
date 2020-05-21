/*
__author__ = 'lawtech'
__date__ = '2020/05/21 10:38 上午'
*/

package _448

import "math"

// hash map
func findDisappearedNumbers(nums []int) []int {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num] = 1
	}

	var res []int
	for num := 1; num < len(nums)+1; num++ {
		if _, ok := hash[num]; !ok {
			res = append(res, num)
		}
	}

	return res
}

// in place
func findDisappearedNumbers1(nums []int) []int {
	for _, num := range nums {
		newIndex := int(math.Abs(float64(num))) - 1
		if nums[newIndex] > 0 {
			nums[newIndex] *= -1
		}
	}

	var res []int
	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] > 0 {
			res = append(res, i)
		}
	}

	return res
}
