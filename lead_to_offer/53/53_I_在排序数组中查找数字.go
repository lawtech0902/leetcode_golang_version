/*
__author__ = 'lawtech'
__date__ = '2020/08/06 9:43 上午'
*/

package _53

// 两次二分，分别查找左右边界
func search(nums []int, target int) int {
	// 搜索右边界right
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	right := i
	// 若数组中无target，则提前返回
	if j >= 0 && nums[j] != target {
		return 0
	}

	// 搜索左边界left
	i, j = 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	left := j
	return right - left - 1
}

// 优化，包装二分
func search1(nums []int, target int) int {
	return helper(nums, target) - helper(nums, target-1)
}

func helper(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return i
}
