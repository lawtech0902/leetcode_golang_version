/*
__author__ = 'lawtech'
__date__ = '2020/1/17 9:56 上午'
*/

package _73

// func setZeroes(matrix [][]int) {
// 	// 暴力
// 	rowNum := len(matrix)
// 	colNum := len(matrix[0])
// 	row := make([]bool, rowNum)
// 	col := make([]bool, colNum)
//
// 	for i := 0; i < rowNum; i++ {
// 		for j := 0; j < colNum; j++ {
// 			if matrix[i][j] == 0 {
// 				row[i] = true
// 				col[j] = true
// 			}
// 		}
// 	}
//
// 	for i := 0; i < rowNum; i++ {
// 		for j := 0; j < colNum; j++ {
// 			if row[i] || col[j] {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
// }

func setZeroes(matrix [][]int) {
	// O(1)空间, 效率还不如暴力。。。。
	rowNum := len(matrix)
	colNum := len(matrix[0])
	isCol := false

	for i := 0; i < rowNum; i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}

		for j := 1; j < colNum; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < rowNum; i++ {
		for j := 1; j < colNum; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < colNum; j++ {
			matrix[0][j] = 0
		}
	}

	if isCol {
		for i := 0; i < rowNum; i++ {
			matrix[i][0] = 0
		}
	}
}
