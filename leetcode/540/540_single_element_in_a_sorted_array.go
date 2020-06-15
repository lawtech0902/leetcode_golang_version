/*
__author__ = 'lawtech'
__date__ = '2020/06/15 11:27 上午'
*/

package _540

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if mid%2 == 1 {
			mid--
		}

		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}

	return nums[low]
}
