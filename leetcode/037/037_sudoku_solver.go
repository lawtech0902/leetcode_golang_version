/*
__author__ = 'lawtech'
__date__ = '2020/1/13 2:09 下午'
*/

package _37

func solveSudoku(board [][]byte) {
	_ = dfs(board)
}

func dfs(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for _, val := range "123456789" {
					board[i][j] = byte(val)
					if isValid(i, j, board) && dfs(board) {
						return true
					}

					board[i][j] = '.'
				}

				return false
			}
		}
	}

	return true
}

func isValid(x, y int, board [][]byte) bool {
	tmp := board[x][y]
	board[x][y] = 'D'

	for i := 0; i < 9; i++ {
		if board[i][y] == tmp {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if board[x][j] == tmp {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[(x/3)*3+i][(y/3)*3+j] == tmp {
				return false
			}
		}
	}

	board[x][y] = tmp
	return true
}
