/*
__author__ = 'lawtech'
__date__ = '2020/05/06 2:07 下午'
*/

package _354

import "sort"

type Dolls [][]int

func (d Dolls) Len() int {
	return len(d)
}

func (d Dolls) Less(i, j int) bool {
	if d[i][0] < d[j][0] {
		return true
	} else if d[i][0] > d[j][0] {
		return false
	}

	return d[i][1] < d[j][1]
}

func (d Dolls) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 || len(envelopes) == 1 {
		return len(envelopes)
	}

	// sort by w first
	dolls := Dolls(envelopes)
	sort.Sort(dolls)

	// LIS by h
	n := len(dolls)
	m := make([]int, n)
	res := 1
	for i := 0; i < n; i++ {
		m[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if dolls[j][0] < dolls[i][0] && dolls[j][1] < dolls[i][1] {
				if m[j]+1 > m[i] {
					m[i] = m[j] + 1 //memorize
					if m[i] > res {
						res = m[i] //save the max to return
					}
				}
			}
		}
	}
	return res
}
