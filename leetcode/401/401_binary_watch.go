/*
__author__ = 'lawtech'
__date__ = '2020/05/13 9:55 上午'
*/

package _401

import "fmt"

func readBinaryWatch(num int) []string {
	var res []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if helper(h)+helper(m) == num {
				s := fmt.Sprintf("%d:%02d", h, m)
				res = append(res, s)
			}
		}
	}

	return res
}

func helper(n int) int {
	res := 0
	for n > 0 {
		res += n % 2
		n /= 2
	}

	return res
}
