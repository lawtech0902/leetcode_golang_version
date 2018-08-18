package _27

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午8:51'
*/

func removeElement(nums []int, val int) int {
	j := 0
	for _, num := range nums {
		if num == val {
			continue
		} else {
			nums[j] = num
			j++
		}
	}
	return j
}
