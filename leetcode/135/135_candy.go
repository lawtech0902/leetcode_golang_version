/*
__author__ = 'lawtech'
__date__ = '2020/03/05 5:02 下午'
*/

package _135

func candy(ratings []int) int {
	if ratings == nil || len(ratings) == 0 {
		return 0
	}

	size := len(ratings)
	reward := make([]int, size)

	// 初始化每个人所分糖果，并比较左值
	reward[0] = 1
	res := 1
	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			reward[i] = reward[i-1] + 1
		} else {
			reward[i] = 1
		}

		res += reward[i]
	}

	// 比较右值，使得分最高者糖果最多
	for i := size - 2; i >= 0; i-- {
		diff := reward[i+1] + 1 - reward[i]
		if ratings[i] > ratings[i+1] && diff > 0 {
			reward[i] += diff
			res += diff
		}
	}

	return res
}
