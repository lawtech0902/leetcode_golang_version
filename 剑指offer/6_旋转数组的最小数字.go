package 剑指offer

/*
__author__ = 'lawtech'
__date__ = '2018/8/26 下午9:11'
*/

func minNumberInRotateArray(rotateArray []int) int {
	size := len(rotateArray)
	if size == 0 {
		return 0
	}

	low, high := 0, size-1
	for rotateArray[low] >= rotateArray[high] {
		if high-low == 1 {
			return rotateArray[high]
		}

		mid := low + (high-low)/2
		if rotateArray[mid] >= rotateArray[low] {
			low = mid
		} else {
			high = mid
		}
	}

	return rotateArray[low]
}
