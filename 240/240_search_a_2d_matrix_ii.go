package _240

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午9:21'
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

	i, j := m-1, 0
	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			j ++
		} else {
			i --
		}
	}

	return false
}
