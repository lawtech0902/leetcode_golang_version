/*
__author__ = 'lawtech'
__date__ = '2020/1/10 4:59 下午'
*/

package _36

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]int)
	columns := make(map[int]map[byte]int)
	boxes := make(map[int]map[byte]int)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				boxIndex := (i/3)*3 + j/3

				if _, ok := rows[i]; !ok {
					rows[i] = make(map[byte]int)
				}

				if _, ok := columns[j]; !ok {
					columns[j] = make(map[byte]int)
				}

				if _, ok := boxes[boxIndex]; !ok {
					boxes[boxIndex] = make(map[byte]int)
				}

				if _, ok := rows[i][num]; ok {
					return false
				} else {
					rows[i][num] = 1
				}

				if _, ok := columns[j][num]; ok {
					return false
				} else {
					columns[j][num] = 1
				}

				if _, ok := boxes[boxIndex][num]; ok {
					return false
				} else {
					boxes[boxIndex][num] = 1
				}
			}
		}
	}

	return true
}
