/*
__author__ = 'lawtech'
__date__ = '2020/04/16 2:48 下午'
*/

package _283

func moveZeroes(nums []int) {
	count := 0
	for _, num := range nums {
		if num != 0 {
			nums[count] = num
			count++
		}
	}

	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}
