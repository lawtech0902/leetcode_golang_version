/*
__author__ = 'lawtech'
__date__ = '2020/06/15 1:48 下午'
*/

package _542

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	n, m := len(matrix), len(matrix[0])
	coordinates := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				coordinates = append(coordinates, [2]int{i, j})
				continue
			}
			matrix[i][j] = 1<<31 - 1
		}
	}

	ways := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(coordinates) > 0 {
		cur := coordinates[0]
		curVal := matrix[cur[0]][cur[1]]
		coordinates = coordinates[1:]

		for _, way := range ways {
			i := cur[0] + way[0]
			j := cur[1] + way[1]

			if i < 0 || j < 0 || i == n || j == m || matrix[i][j] <= curVal+1 {
				continue
			}

			matrix[i][j] = curVal + 1
			coordinates = append(coordinates, [2]int{i, j})
		}
	}

	return matrix
}
