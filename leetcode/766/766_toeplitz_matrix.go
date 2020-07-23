/*
__author__ = 'lawtech'
__date__ = '2020/07/23 9:47 上午'
*/

package _766

// 检查左上邻居
func isToeplitzMatrix(matrix [][]int) bool {
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[r][c] != matrix[r-1][c-1] {
				return false
			}
		}
	}

	return true
}

// map存储 暴力
func isToeplitzMatrix1(matrix [][]int) bool {
	groups := make(map[int]int)
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if val, ok := groups[r-c]; !ok {
				groups[r-c] = matrix[r][c]
			} else if val != matrix[r][c] {
				return false
			}
		}
	}

	return true
}
