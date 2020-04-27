/*
__author__ = 'lawtech'
__date__ = '2020/04/26 2:01 下午'
*/

package _332

import "sort"

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string, len(tickets)+1)
	var res []string

	for _, t := range tickets {
		m[t[0]] = append(m[t[0]], t[1])
	}

	for k := range m {
		sort.Strings(m[k])
	}

	dfs("JFK", m, &res)

	// reverse res arr
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

func dfs(start string, m map[string][]string, res *[]string) {
	for len(m[start]) > 0 {
		cur := m[start][0]
		m[start] = m[start][1:]
		dfs(cur, m, res)
	}

	*res = append(*res, start)
}
