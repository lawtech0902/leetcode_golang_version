/*
__author__ = 'lawtech'
__date__ = '2020/05/14 3:20 下午'
*/

package _417

// dfs
func pacificAtlantic(matrix [][]int) [][]int {
	if matrix == nil || (matrix != nil && len(matrix) == 0) {
		return nil
	}

	var res [][]int
	m, n := len(matrix), len(matrix[0])
	pac := make([][]bool, m)
	atl := make([][]bool, m)

	for i := range pac {
		pac[i] = make([]bool, n)
		atl[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, matrix, pac)
	}

	for j := 0; j < n; j++ {
		dfs(0, j, matrix, pac)
	}

	for i := 0; i < m; i++ {
		dfs(i, n-1, matrix, atl)
	}

	for j := 0; j < n; j++ {
		dfs(m-1, j, matrix, atl)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pac[i][j] && atl[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func dfs(i, j int, matrix [][]int, visited [][]bool) {
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	m, n := len(matrix), len(matrix[0])
	visited[i][j] = true

	for k := 0; k < 4; k++ {
		x, y := i+dirs[k][0], j+dirs[k][1]
		if x < 0 || x >= m || y < 0 || y >= n {
			continue
		}

		if visited[x][y] {
			continue
		}

		if matrix[x][y] < matrix[i][j] {
			continue
		}

		dfs(x, y, matrix, visited)
	}
}
