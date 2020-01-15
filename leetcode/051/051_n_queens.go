/*
__author__ = 'lawtech'
__date__ = '2020/1/15 4:13 下午'
*/

package _51

func solveNQueens(n int) (result [][]string) {
	board := makeBoard(n)
	place(n, 0, board, &result)
	return
}

func place(n, row int, board [][]byte, result *[][]string) {
	if row == n {
		cp := make([]string, n)
		for i := range board {
			cp[i] = string(board[i])
		}
		*result = append(*result, cp)
		return
	}
	for i := range board[row] {
		// 垂直是否有
		for j := 0; j < n; j++ {
			if board[j][i] != '.' {
				goto next
			}
		}
		// 斜对角 row， i
		for a, b := row-1, i-1; a >= 0 && b >= 0; {
			if board[a][b] != '.' {
				goto next
			}
			a--
			b--
		}
		for a, b := row-1, i+1; a >= 0 && b < n; {
			if board[a][b] != '.' {
				goto next
			}
			a--
			b++
		}
		for a, b := row+1, i-1; a < n && b >= 0; {
			if board[a][b] != '.' {
				goto next
			}
			a++
			b--
		}
		for a, b := row+1, i+1; a < n && b < n; {
			if board[a][b] != '.' {
				goto next
			}
			a++
			b++
		}
		// 放置-》下一步
		board[row][i] = 'Q'
		place(n, row+1, board, result)
		board[row][i] = '.'
	next:
		continue
	}
}

func makeBoard(n int) [][]byte {
	result := make([][]byte, n)
	for i := range result {
		s := make([]byte, n)
		for i := 0; i < n; i++ {
			s[i] = '.'
		}
		result[i] = s
	}
	return result
}
