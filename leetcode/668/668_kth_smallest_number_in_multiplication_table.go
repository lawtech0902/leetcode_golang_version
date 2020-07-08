/*
__author__ = 'lawtech'
__date__ = '2020/07/08 2:01 下午'
*/

package _668

func findKthNumber(m, n, k int) int {
	left, right := 1, m*n
	for left < right {
		mid := left + (right-left)/2
		if kthCount(m, n, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func kthCount(m, n, mid int) int {
	res := 0
	row, col := 1, m
	for col > 0 && row <= n {
		if col*row <= mid {
			res += col
			row++
		} else {
			col--
		}
	}

	return res
}
