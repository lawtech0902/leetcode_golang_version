/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午10:00'
*/

package _169

func majorityElement(nums []int) int {
	candidate, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
			count = 1
		} else if num == candidate {
			count += 1
		} else {
			count -= 1
		}
	}

	return candidate
}
