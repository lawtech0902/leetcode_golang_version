/*
__author__ = 'lawtech'
__date__ = '2020/06/03 9:53 上午'
*/

package _498

func findDiagonalOrder(matrix [][]int) []int {
	var res []int

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	n, m := len(matrix), len(matrix[0])

	for d := 0; d < n+m-1; d++ {
		var (
			intermediate []int
			r, c         int
		)

		if d < m {
			r, c = 0, d
		} else {
			r, c = d-m+1, m-1
		}

		for r < n && c > -1 {
			intermediate = append(intermediate, matrix[r][c])
			r++
			c--
		}

		if d%2 == 0 {
			intermediate = reverseIntSlice(intermediate)
		}

		res = append(res, intermediate...)
	}

	return res
}

func reverseIntSlice(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
