/*
__author__ = 'lawtech'
__date__ = '2020/03/12 6:28 下午'
*/

package _162

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
