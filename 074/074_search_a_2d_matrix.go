package _74

/*
__author__ = 'lawtech'
__date__ = '2018/8/24 下午5:21'
*/

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	i, j := 0, n-1

	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j -= 1
		} else {
			i += 1
		}
	}

	return false
}
