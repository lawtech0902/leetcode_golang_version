/*
__author__ = 'lawtech'
__date__ = '2020/04/26 9:45 上午'
*/

package _327

func countRangeSum(nums []int, lower int, upper int) int {
	rst := 0
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		tmp := 0
		for j := i; j < lenNums; j++ {
			if i == j {
				tmp = nums[i]
			} else {
				tmp = tmp + nums[j]
			}
			if tmp <= upper && tmp >= lower {
				rst++
			}

		}
	}

	return rst
}
