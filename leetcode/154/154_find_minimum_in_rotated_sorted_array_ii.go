/*
__author__ = 'lawtech'
__date__ = '2020/03/11 2:29 下午'
*/

package _154

func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := int(uint(i+j) >> 1)

		if nums[mid] > nums[j] {
			i = mid + 1
		} else if nums[mid] < nums[j] {
			j = mid
		} else {
			j--
		}
	}
	return nums[i]
}
