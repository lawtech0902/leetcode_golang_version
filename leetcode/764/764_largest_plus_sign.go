/*
__author__ = 'lawtech'
__date__ = '2020/07/22 3:52 下午'
*/

package _764

func orderOfLargestPlusSign(N int, mines [][]int) int {
	grid := make([][]int, N)
	for i := range grid {
		grid[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			grid[i][j] = N
		}
	}

	for _, v := range mines {
		grid[v[0]][v[1]] = 0
	}

	for i := 0; i < N; i++ {
		left, right, up, down := 0, 0, 0, 0
		for j, k := 0, N-1; j < N; j, k = j+1, k-1 {
			if grid[i][j] == 0 {
				left = 0
			} else {
				left++
			}

			if left < grid[i][j] {
				grid[i][j] = left
			}

			if grid[i][k] == 0 {
				right = 0
			} else {
				right += 1
			}

			if right < grid[i][k] {
				grid[i][k] = right
			}

			if grid[j][i] == 0 {
				up = 0
			} else {
				up += 1
			}

			if up < grid[j][i] {
				grid[j][i] = up
			}

			if grid[k][i] == 0 {
				down = 0
			} else {
				down += 1
			}

			if down < grid[k][i] {
				grid[k][i] = down
			}
		}
	}

	res := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] > res {
				res = grid[i][j]
			}
		}
	}

	return res
}
