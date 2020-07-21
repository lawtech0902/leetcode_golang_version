/*
__author__ = 'lawtech'
__date__ = '2020/07/21 2:11 下午'
*/

package _747

func dominantIndex(nums []int) int {
	highest, secHighest := 0, 0
	idx := 0
	for k, v := range nums {
		if v > highest {
			secHighest = highest
			highest = v
			idx = k
		} else if v > secHighest {
			secHighest = v
		}
	}

	if 2*secHighest <= highest {
		return idx
	}

	return -1
}
