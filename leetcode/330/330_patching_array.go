/*
__author__ = 'lawtech'
__date__ = '2020/04/26 11:25 上午'
*/

package _330

// hard == not necessary to do
func minPatches(nums []int, n int) int {
	miss := 1 // max value can reached
	added := 0
	i := 0
	for miss <= n {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			added++
		}
	}

	return added
}
