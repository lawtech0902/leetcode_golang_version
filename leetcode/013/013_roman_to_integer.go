/*
__author__ = 'lawtech'
__date__ = '2018/8/18 上午12:47'
*/

package _13

func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		tmp := m[s[i]]

		sign := 1
		if tmp < last {
			// 小数在大数的左边，要减去小数
			sign = -1
		}

		res += sign * tmp
		last = tmp
	}

	return res
}
