/*
__author__ = 'lawtech'
__date__ = '2020/04/08 10:28 上午'
*/

package _227

import (
	"unicode"
)

func calculate(s string) int {
	stack := make([]int, 0)
	num := 0
	sign := '+'

	for i, r := range s {
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')
			if i != len(s)-1 {
				continue
			}
		}

		if r == ' ' && i != len(s)-1 {
			continue
		}

		switch sign {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			newNum := stack[len(stack)-1] * num
			stack[len(stack)-1] = newNum
		case '/':
			newNum := stack[len(stack)-1] / num
			stack[len(stack)-1] = newNum
		}

		num = 0
		sign = r
	}

	res := 0
	for _, el := range stack {
		res += el
	}

	return res
}
