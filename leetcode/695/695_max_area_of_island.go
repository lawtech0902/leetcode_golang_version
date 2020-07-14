/*
__author__ = 'lawtech'
__date__ = '2020/07/14 10:12 上午'
*/

package _695

// dfs
func maxAreaOfIsland(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			area := 0
			dfs(i, j, grid, &area)
			if area > res {
				res = area
			}
		}
	}

	return res
}

func dfs(i, j int, grid [][]int, area *int) {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == 0 {
		return
	}

	*area++
	grid[i][j] = 0
	dfs(i+1, j, grid, area)
	dfs(i-1, j, grid, area)
	dfs(i, j+1, grid, area)
	dfs(i, j-1, grid, area)
}
