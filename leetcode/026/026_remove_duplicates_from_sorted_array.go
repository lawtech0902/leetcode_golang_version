/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午8:36'
*/

package _26

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	cur := 0
	for i := 1; i < size; i++ {
		if nums[i] != nums[cur] {
			cur += 1
			nums[cur] = nums[i]
		}
	}

	return cur + 1
}
