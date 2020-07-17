/*
__author__ = 'lawtech'
__date__ = '2020/07/17 10:55 上午'
*/

package _724

func pivotIndex(nums []int) int {
	s := sumSlice(nums)
	leftSum := 0
	for i, num := range nums {
		if leftSum == s-leftSum-num {
			return i
		}

		leftSum += num
	}

	return -1
}

func sumSlice(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}
