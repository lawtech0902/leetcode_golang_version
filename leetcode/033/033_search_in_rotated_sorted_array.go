/*
__author__ = 'lawtech'
__date__ = '2020/1/10 3:56 下午'
*/

package _33

func search(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1

	for left < right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
