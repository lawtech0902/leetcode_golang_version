/*
__author__ = 'lawtech'
__date__ = '2020/1/10 4:33 下午'
*/

package _34

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			res := []int{0, 0}
			if nums[left] == target {
				res[0] = left
			}
			if nums[right] == target {
				res[1] = right
			}

			for i := mid; i < right+1; i++ {
				if nums[i] != target {
					res[1] = i - 1
					break
				}
			}

			for i := mid; i > left-1; i-- {
				if nums[i] != target {
					res[0] = i + 1
					break
				}
			}

			return res
		}
	}

	return []int{-1, -1}
}
