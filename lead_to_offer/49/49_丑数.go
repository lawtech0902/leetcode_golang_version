/*
__author__ = 'lawtech'
__date__ = '2020/08/05 2:05 下午'
*/

package _49

import "math"

func nthUglyNumber(n int) int {
	q := make([]int, n)
	q[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < len(q); i++ {
		m2, m3, m5 := q[i2]*2, q[i3]*3, q[i5]*5
		m := min(m2, m3, m5)
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

func min(values ...int) int {
	minValue := math.MaxInt32
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}
