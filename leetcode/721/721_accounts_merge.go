/*
__author__ = 'lawtech'
__date__ = '2020/07/16 3:32 下午'
*/

package _721

import "sort"

// dfs
func accountsMerge(accounts [][]string) [][]string {
	emailToName := make(map[string]string)
	graph := make(map[string][]string)
	for _, v := range accounts {
		name := v[0]
		for _, email := range v[1:] {
			graph[v[1]] = append(graph[v[1]], email)
			graph[email] = append(graph[email], v[1])
			emailToName[email] = name
		}
	}

	usedEmail := make(map[string]bool)
	var res [][]string
	for email := range graph {
		if usedEmail[email] {
			continue
		}

		usedEmail[email] = true
		queue := []string{email}
		var component []string
		for len(queue) != 0 {
			firstEmail := queue[0]
			queue = queue[1:]
			component = append(component, firstEmail)
			for _, neighbour := range graph[firstEmail] {
				if usedEmail[neighbour] {
					continue
				}

				usedEmail[neighbour] = true
				queue = append(queue, neighbour)
			}
		}

		sort.Strings(component)
		component = append([]string{emailToName[email]}, component...)
		res = append(res, component)
	}

	return res
}
