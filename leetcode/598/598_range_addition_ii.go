/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-23 14:42:30
 */

package _598

func maxCount(m, n int, ops [][]int) int {
	min1, min2 := m, n
	for _, op := range ops {
		min1 := min(min1, op[0])
		min2 := min(min2, op[1])
	}

	return min1 * min2
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
