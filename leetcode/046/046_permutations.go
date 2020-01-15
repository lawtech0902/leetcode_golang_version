/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:34'
*/

package _46

func permute(nums []int) [][]int {
	size := len(nums)
	if size <= 1 {
		return [][]int{nums}
	}
	var res [][]int
	for i := range nums {
		var numsCopy = make([]int, len(nums))
		copy(numsCopy, nums) // 很重要
		for _, j := range permute(append(numsCopy[:i], numsCopy[i+1:]...)) {
			res = append(res, append([]int{nums[i]}, j...))
		}
	}
	return res
}
