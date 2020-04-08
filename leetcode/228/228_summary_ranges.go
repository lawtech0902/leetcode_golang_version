/*
__author__ = 'lawtech'
__date__ = '2020/04/08 10:45 上午'
*/

package _228

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var (
		res  []string
		head int
	)

	for i := range nums {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}

		if head == i {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			tmp := strconv.Itoa(nums[head]) + "->" + strconv.Itoa(nums[i])
			res = append(res, tmp)
		}

		head = i + 1
	}

	return res
}
