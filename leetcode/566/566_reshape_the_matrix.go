/*
__author__ = 'lawtech'
__date__ = '2020/06/18 9:47 上午'
*/

package _566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if r*c != len(nums)*len(nums[0]) {
		return nums
	}

	arr := make([]int, 0)
	for _, v := range nums {
		for _, a := range v {
			arr = append(arr, a)
			if len(arr) == c {
				nums = append(nums, arr)
				arr = []int{}
			}
		}
	}

	return nums[len(nums)-r:]
}
