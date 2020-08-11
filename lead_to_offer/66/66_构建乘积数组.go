/*
__author__ = 'lawtech'
__date__ = '2020/08/11 2:20 下午'
*/

package _66

func constructArr(a []int) []int {
	n := len(a)
	res := make([]int, n)
	left := 1
	for i := 0; i < n; i++ {
		// 把左边的数的乘积存入res
		res[i] = left
		left *= a[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		// 把左边的数的乘积乘以右边的数的乘积，并存入res
		res[i] *= right
		// 更新右边数的乘积
		right *= a[i]
	}

	return res
}
