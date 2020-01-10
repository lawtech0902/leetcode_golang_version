/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午8:51'
*/

package _27

func removeElement(nums []int, val int) int {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		} else {
			nums[cur] = nums[i]
			cur += 1
		}
	}

	return cur
}
