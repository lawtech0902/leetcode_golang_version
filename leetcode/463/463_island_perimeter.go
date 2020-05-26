/*
__author__ = 'lawtech'
__date__ = '2020/05/26 10:33 上午'
*/

package _463

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					res -= 2
				}

				if j-1 >= 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}
