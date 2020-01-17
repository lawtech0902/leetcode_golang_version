/*
__author__ = 'lawtech'
__date__ = '2020/1/17 1:57 下午'
*/

package _80

func removeDuplicates(nums []int) int {
	i := 0
	for _, num := range nums {
		if i < 2 || num > nums[i-2] {
			nums[i] = num
			i += 1
		}
	}

	return i
}
