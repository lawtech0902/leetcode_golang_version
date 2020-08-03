/*
__author__ = 'lawtech'
__date__ = '2020/08/03 1:42 下午'
*/

package _39

// 摩尔投票法
func majorityElement(nums []int) int {
	x, votes := 0, 0
	for _, num := range nums {
		if votes == 0 {
			x = num
		}

		if num == x {
			votes++
		} else {
			votes--
		}
	}

	return x
}
