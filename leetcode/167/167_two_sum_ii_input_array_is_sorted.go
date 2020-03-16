/*
__author__ = 'lawtech'
__date__ = '2018/8/24 下午5:57'
*/

package _167

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		s := numbers[l] + numbers[r]
		if s == target {
			return []int{l + 1, r + 1}
		} else if s < target {
			l++
		} else {
			r--
		}
	}

	return []int{}
}
