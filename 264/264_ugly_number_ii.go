package _264

import (
	"go_projects/leetcode_golang_version"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午11:22'
*/

func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}

	q := []int{1}
	i2, i3, i5 := 0, 0, 0

	for len(q) < n {
		m2, m3, m5 := q[i2]*2, q[i3]*3, q[i5]*5
		m, _ := leetcode_golang_version.MinMax([]int{m2, m3, m5})
		switch m {
		case 2:
			i2 += 1
		case 3:
			i3 += 1
		case 5:
			i5 += 1
		default:
		}
		q = append(q, m)
	}

	return q[len(q)-1]
}
