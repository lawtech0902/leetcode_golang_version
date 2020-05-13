/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午10:28'
*/

package _406

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// sort input based on h first then k
	// if h is equal, place element with bigger k value first
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}

		return people[i][0] < people[j][0]
	})

	// init
	res := make([][]int, len(people))
	index := make([]int, len(people))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, 2)
		index[i] = i
	}

	// place people at index k
	for _, p := range people {
		i := index[p[1]]
		res[i][0], res[i][1] = p[0], p[1]
		index = append(index[0:p[1]], index[p[1]+1:]...)
	}

	return res
}
