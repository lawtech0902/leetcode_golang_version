package _315

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午11:46'
*/

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				res[i]++
			}
		}
	}

	return res
}
