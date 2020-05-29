/*
__author__ = 'lawtech'
__date__ = '2020/05/29 1:31 下午'
*/

package _485

func findMaxConsecutiveOnes(nums []int) int {
	count, maxCount := 0, 0
	for _, num := range nums {
		if num == 1 {
			count += 1
		} else {
			maxCount = max(maxCount, count)
			count = 0
		}
	}

	return max(maxCount, count)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
