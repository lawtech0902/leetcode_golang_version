/*
__author__ = 'lawtech'
__date__ = '2020/03/25 1:24 下午'
*/

package _210

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	state := make([]int, numCourses)
	var ans []int

	for i := 0; i < numCourses; i++ {
		if DFS(graph, i, state, &ans) == false {
			return []int{}
		}
	}

	ans = reverse(ans)
	return ans
}

func DFS(graph map[int][]int, cur int, state []int, ans *[]int) bool {
	if state[cur] == 1 {
		return true
	}

	if state[cur] == 2 {
		return false
	}

	state[cur] = 2
	for _, v := range graph[cur] {
		if DFS(graph, v, state, ans) == false {
			return false
		}
	}
	state[cur] = 1

	*ans = append(*ans, cur)

	return true
}

func reverse(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
