/*
__author__ = 'lawtech'
__date__ = '2020/05/08 2:20 下午'
*/

package _376

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	flag := 0 // -1 need a smaller number, 1 need a bigger number, 0 need a smaller/bigger number
	last, maxLength := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if flag == 0 {
			if nums[i] > last {
				flag = -1
				maxLength++
			} else if nums[i] < last {
				flag = 1
				maxLength++
			}
		} else if flag == 1 {
			if nums[i] > last {
				flag = -1
				maxLength++
			}
		} else {
			if nums[i] < last {
				flag = 1
				maxLength++
			}
		}

		last = nums[i]
	}

	return maxLength
}
