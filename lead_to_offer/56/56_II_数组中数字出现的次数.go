/*
__author__ = 'lawtech'
__date__ = '2020/08/07 10:48 上午'
*/

package _56

// hash，别的不会
func singleNumber(nums []int) int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}

	for i, v := range hash {
		if v == 1 {
			return i
		}
	}

	return 0
}
