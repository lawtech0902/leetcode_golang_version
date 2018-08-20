package _136

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午11:11'
*/

func singleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		res ^= num
	}

	return res
}
