/*
__author__ = 'lawtech'
__date__ = '2020/03/30 2:12 下午'
*/

package _217

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, n := range nums {
		if exist := set[n]; exist {
			return true
		}

		set[n] = true
	}

	return false
}
