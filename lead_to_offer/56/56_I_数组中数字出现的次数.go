/*
__author__ = 'lawtech'
__date__ = '2020/08/06 2:31 下午'
*/

package _56

func singleNumbers(nums []int) []int {
	var a int
	for _, v := range nums {
		a ^= v
	}

	mask := a & (-a)
	res := make([]int, 2)
	for _, v := range nums {
		if v&mask == 0 { // 分组
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}

	return res
}
