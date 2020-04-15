/*
__author__ = 'lawtech'
__date__ = '2020/04/15 1:29 下午'
*/

package _278

// skip function body
func isBadVersion(n int) bool {
	if n == 1 {
		return false
	}

	return true
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
