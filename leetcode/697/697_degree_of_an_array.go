/*
__author__ = 'lawtech'
__date__ = '2020/07/14 11:33 上午'
*/

package _697

func findShortestSubArray(nums []int) int {
	left, right, count := make(map[int]int), make(map[int]int), make(map[int]int)
	for i, num := range nums {
		if _, ok := left[num]; !ok {
			left[num] = i
		}

		right[num] = i
		count[num]++
	}

	res := len(nums)
	degree := 0
	for _, v := range count {
		if v > degree {
			degree = v
		}
	}

	for k := range count {
		if count[k] == degree {
			if right[k]-left[k]+1 < res {
				res = right[k] - left[k] + 1
			}
		}
	}

	return res
}
