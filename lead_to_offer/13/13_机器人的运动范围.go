/*
__author__ = 'lawtech'
__date__ = '2020/07/28 11:23 上午'
*/

package _13

// dfs
func movingCount(m, n, k int) int {
	var dfs func(i, j, si, sj int) int

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dfs = func(i, j, si, sj int) int {
		if i >= m || j >= n || si+sj > k || visited[i][j] {
			return 0
		}

		visited[i][j] = true
		return 1 + dfs(i+1, j, sumX(i, si), sj) + dfs(i, j+1, si, sumX(j, sj))
	}

	return dfs(0, 0, 0, 0)
}

// bfs
func movingCount1(m, n, k int) int {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	res := 0
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0, 0, 0})
	for len(queue) > 0 {
		i, j, si, sj := queue[0][0], queue[0][1], queue[0][2], queue[0][3]
		// pop
		queue = queue[1:]
		if i >= m || j >= n || si+sj > k || visited[i][j] {
			continue
		}

		visited[i][j] = true
		res++
		queue = append(queue, []int{i + 1, j, sumX(i, si), sj})
		queue = append(queue, []int{i, j + 1, si, sumX(j, sj)})
	}

	return res
}

// 数位和递增
func sumX(x, sx int) int {
	if (x+1)%10 != 0 {
		return sx + 1
	} else {
		return sx - 8
	}
}
