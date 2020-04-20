/*
__author__ = 'lawtech'
__date__ = '2020/04/17 2:21 下午'
*/

package _289

/*
规则：1. 原来是活的，周围有2-3个活的，成为活的
	 2. 原来是死的，周围有3个活的，成为活的
	 3. 其他都是死了
换个角度，遍历的时候，如果这个格子里是活的，我就让它去“影响”周围的八个格子。
这样一来，大批原来是死了的格子就省了很多检查的时间。
怎么“影响”？简单，我给被影响的格子里的数字加10。
这样一来，个位存着这格子原来的状态，十位就存着它周围有多少个活格子了。
比如遍历完了之后一个格子里是41，那就表示它原来自己是1，然后被周围的四个活格子加了四个10，于是周围有四个活细胞。
*/
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	affect := func(x, y int) {
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if i < 0 || i >= m || j < 0 || j >= n || (i == x && j == y) {
					continue
				}

				board[i][j] += 10
			}
		}
	}

	calculate := func(x, y int) {
		value := board[x][y]
		if value/10 == 3 {
			board[x][y] = 1
		} else if value/10 == 2 && value%10 == 1 {
			board[x][y] = 1
		} else {
			board[x][y] = 0
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j]%10 == 1 {
				affect(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			calculate(i, j)
		}
	}
}
