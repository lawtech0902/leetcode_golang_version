/*
__author__ = 'lawtech'
__date__ = '2020/05/07 10:29 上午'
*/

package _367

// binary search O(logn) O(1)
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	l, r := 0, num/2
	for l < r {
		m := l + (l-r)/2
		if m*m < num {
			l = m + 1
		} else if m*m > num {
			r = m
		} else {
			return true
		}
	}

	return false
}

// newton iterative method O(logn) O(1)
func isPerfectSquare1(num int) bool {
	if num < 2 {
		return true
	}

	x := num / 2
	for x*x > num {
		x = (x + num/x) / 2
	}

	return x*x == num
}
