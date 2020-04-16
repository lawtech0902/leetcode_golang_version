/*
__author__ = 'lawtech'
__date__ = '2020/04/16 3:38 下午'
*/

package _287

// floyd cycle
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for ok := true; ok; ok = slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}
