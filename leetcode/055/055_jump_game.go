/*
__author__ = 'lawtech'
__date__ = '2020/1/15 9:55 上午'
*/

package _55

func canJump(nums []int) bool {
	last := len(nums) - 1

	for i := len(nums) - 2; i > -1; i-- {
		if nums[i] >= last-i {
			last = i
		}
	}

	return last == 0
}
