/*
__author__ = 'lawtech'
__date__ = '2020/07/21 10:41 上午'
*/

package _744

// 二分
func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)
	for low < high {
		mid := low + (high-low)/2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return letters[low%len(letters)]
}
