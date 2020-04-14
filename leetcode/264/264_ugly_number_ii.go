/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午11:22'
*/

package _264

import "math"

func nthUglyNumber(n int) int {
	q := make([]int, n)
	q[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < len(q); i++ {
		m2, m3, m5 := q[i2]*2, q[i3]*3, q[i5]*5
		m := Min(m2, m3, m5)
		q[i] = m

		if m == m2 {
			i2++
		}

		if m == m3 {
			i3++
		}

		if m == m5 {
			i5++
		}
	}

	return q[n-1]
}

func Min(args ...int) int {
	min := math.MaxInt32
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}
