/*
__author__ = 'lawtech'
__date__ = '2020/08/06 10:32 上午'
*/

package _53

// O(n)
func missingNumber(nums []int) int {
	if nums[0] == 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}

// 二分 O(nlogn)
func missingNumber1(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return i
}
