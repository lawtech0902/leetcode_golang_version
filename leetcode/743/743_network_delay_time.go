/*
__author__ = 'lawtech'
__date__ = '2020/07/21 9:49 上午'
*/

package _743

// Dijkstra
func networkDelayTime(times [][]int, N, K int) int {
	graphs, values := initialize(times, N)
	return Dijkstra(K, N, &graphs, &values)
}

func initialize(times [][]int, N int) ([][]int, [][]int) {
	graphs := make([][]int, N+1)
	values := make([][]int, N+1)
	for _, time := range times {
		graphs[time[0]] = append(graphs[time[0]], time[1])
		values[time[0]] = append(values[time[0]], time[2])
	}

	return graphs, values
}

func Dijkstra(K, N int, graphs, values *[][]int) int {
	queue := []int{K}
	visited := make([]int, N+1)
	for i := range visited {
		visited[i] = -1
	}

	visited[K] = 0

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for i, neighbor := range (*graphs)[curr] {
			if visited[neighbor] == -1 || visited[neighbor] > visited[curr]+(*values)[curr][i] {
				visited[neighbor] = visited[curr] + (*values)[curr][i]
				queue = append(queue, neighbor)
			}
		}
	}

	max := -1
	for i := 1; i < N+1; i++ {
		if visited[i] == -1 {
			return -1
		}

		if visited[i] > max {
			max = visited[i]
		}
	}

	return max
}
