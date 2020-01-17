/*
__author__ = 'lawtech'
__date__ = '2020/1/17 1:32 下午'
*/

package _79

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, board, word[1:]) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(x, y int, board [][]byte, word string) bool {

	if len(word) == 0 {
		return true
	}

	// up
	if x > 0 && board[x-1][y] == word[0] {
		tmp := board[x][y]
		board[x][y] = '#'
		if dfs(x-1, y, board, word[1:]) {
			return true
		}
		board[x][y] = tmp
	}

	// down
	if x < len(board)-1 && board[x+1][y] == word[0] {
		tmp := board[x][y]
		board[x][y] = '#'
		if dfs(x+1, y, board, word[1:]) {
			return true
		}
		board[x][y] = tmp
	}

	// left
	if y > 0 && board[x][y-1] == word[0] {
		tmp := board[x][y]
		board[x][y] = '#'
		if dfs(x, y-1, board, word[1:]) {
			return true
		}
		board[x][y] = tmp
	}

	// right
	if y < len(board[0])-1 && board[x][y+1] == word[0] {
		tmp := board[x][y]
		board[x][y] = '#'
		if dfs(x, y+1, board, word[1:]) {
			return true
		}
		board[x][y] = tmp
	}

	return false
}
