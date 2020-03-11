/*
__author__ = 'lawtech'
__date__ = '2020/03/11 2:14 下午'
*/

package _153

func findMin(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	l, r := 0, size-1
	if nums[r] > nums[0] {
		return nums[0]
	}

	for r >= l {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
