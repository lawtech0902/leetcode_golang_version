/*
__author__ = 'lawtech'
__date__ = '2020/05/20 3:35 下午'
*/

package _442

import "math"

// 负号自哈希
func findDuplicates(nums []int) []int {
	var res []int

	for _, num := range nums {
		num = int(math.Abs(float64(num)))
		if nums[num-1] > 0 {
			nums[num-1] *= -1
		} else {
			res = append(res, num)
		}
	}

	return res
}
