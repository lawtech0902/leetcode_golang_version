/*
__author__ = 'lawtech'
__date__ = '2020/1/16 4:09 下午'
*/

package _69

func mySqrt(x int) int {
	// 二分法
	if x <= 1 {
		return x
	}

	low, high := 1, x
	for low <= high {
		mid := low + (high-low)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt > mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}
