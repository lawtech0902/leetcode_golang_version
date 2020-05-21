/*
__author__ = 'lawtech'
__date__ = '2020/05/21 2:31 下午'
*/

package _451

import "strings"

func frequencySort(s string) string {
	// bucket to store rune with the same frequency
	bucket := make([][]rune, len(s)+1)
	// map to store rune's frequency
	m := make(map[rune]int)
	var res string

	for _, r := range s {
		m[r] += 1
	}

	for r, time := range m {
		bucket[time] = append(bucket[time], r)
	}

	for i := len(s); i >= 0; i-- {
		for _, r := range bucket[i] {
			res += strings.Repeat(string(r), i)
		}
	}

	return res
}
