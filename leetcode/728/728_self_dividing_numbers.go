/*
__author__ = 'lawtech'
__date__ = '2020/07/17 2:04 下午'
*/

package _728

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}

	return res
}

func isSelfDividing(n int) bool {
	x := n
	for x > 0 {
		d := x % 10
		x = x / 10
		if d == 0 || n%d > 0 {
			return false
		}
	}

	return true
}
