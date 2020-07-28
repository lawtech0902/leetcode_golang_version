/*
__author__ = 'lawtech'
__date__ = '2020/07/28 10:27 上午'
*/

package _11

func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}

	return numbers[low]
}
