/*
__author__ = 'lawtech'
__date__ = '2018/8/24 下午5:21'
*/

package _74

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	j := len(matrix[0]) - 1
	for i := 0; i < len(matrix); i++ {
		for j != 0 && matrix[i][j] > target {
			j -= 1
		}

		if matrix[i][j] == target {
			return true
		}
	}

	return false
}
