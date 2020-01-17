/*
__author__ = 'lawtech'
__date__ = '2020/1/17 2:01 下午'
*/

package _81

func search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] {
			left += 1
			continue
		}

		// 前半部分有序
		if nums[left] < nums[mid] {
			// target 在前半部分
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1 // 否则，去后半部分找
			}
		} else {
			// 后半部分有序
			// target 在后半部分
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1 // 否则，去前半部分找
			}
		}
	}

	return false
}
