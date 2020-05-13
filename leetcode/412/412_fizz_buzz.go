/*
__author__ = 'lawtech'
__date__ = '2020/05/13 3:55 下午'
*/

package _412

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var s string
		if i%3 == 0 {
			s += "Fizz"
		}

		if i%5 == 0 {
			s += "Buzz"
		}

		if s == "" {
			s = strconv.Itoa(i)
		}

		res = append(res, s)
	}

	return res
}
