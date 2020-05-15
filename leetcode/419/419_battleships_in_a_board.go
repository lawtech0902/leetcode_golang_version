/*
__author__ = 'lawtech'
__date__ = '2020/05/15 9:51 上午'
*/

package _419

func countBattleships(board [][]byte) int {
	var count int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				count++
			}
		}
	}

	return count
}
