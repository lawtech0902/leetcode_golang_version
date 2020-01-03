package _455

import "sort"

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午11:59'
*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if g[gi] <= s[si] {
			gi++
		}
		si ++
	}

	return gi
}
