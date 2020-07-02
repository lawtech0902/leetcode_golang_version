/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-02 11:16:44
 */

package _640

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	s := strings.Split(equation, "=")
	if len(s) != 2 {
		return "No solution"
	}

	l, r := s[0], s[1]
	lx, ln := helper(l)
	rx, rn := helper(r)
	x := lx - rx
	n := rn - ln
	if x == 0 {
		if n == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}

	return "x=" + strconv.Itoa(n/x)
}

func helper(equation string) (int, int) {
	equation += "+"
	sign := 1
	current := ""
	coefficient, number := 0, 0
	for _, v := range equation {
		switch v {
		case '+':
			n, _ := strconv.Atoi(current)
			number += sign * n
			sign = 1
			current = ""
		case '-':
			n, _ := strconv.Atoi(current)
			number += sign * n
			sign = -1
			current = ""
		case 'x':
			if current == "" {
				current = "1"
			}

			n, _ := strconv.Atoi(current)
			coefficient += sign * n
			current = "0"
		default:
			current += string(v)
		}
	}

	return coefficient, number
}
