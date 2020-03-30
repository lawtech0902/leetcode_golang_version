/*
__author__ = 'lawtech'
__date__ = '2020/03/30 4:54 下午'
*/

package _219

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	size := len(nums)
	if size == 0 {
		return false
	}

	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			if int(math.Abs(float64(m[num]-i))) <= k {
				return true
			} else {
				if i < size {
					m[num] = i
					continue
				}
			}
		} else {
			m[num] = i
		}
	}

	return false
}
