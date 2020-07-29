/*
__author__ = 'lawtech'
__date__ = '2020/07/29 10:41 上午'
*/

package _17

import (
	"math"
	"strconv"
)

func printNumbers(n int) []int {
	var (
		nine, count int
		dfs         func(x int)
	)

	loop := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	res := make([]int, int(math.Pow(10, float64(n)))-1)
	num := make([]byte, n)
	start := n - 1

	dfs = func(x int) {
		if x == n {
			s := string(num)[start:]
			if s != "0" {
				res[count], _ = strconv.Atoi(s)
				count++
			}

			if n-start == nine {
				start--
			}

			return
		}

		for _, v := range loop {
			if v == '9' {
				nine++
			}

			num[x] = v
			dfs(x + 1)
		}

		nine--
	}

	dfs(0)
	return res
}
