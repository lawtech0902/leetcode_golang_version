/*
__author__ = 'lawtech'
__date__ = '2020/04/13 3:16 下午'
*/

package _241

import "strconv"

func diffWaysToCompute(input string) []int {
	var res []int

	opChars := map[byte]func(int, int) int{
		'+': func(a, b int) int { return a + b },
		'-': func(a, b int) int { return a - b },
		'*': func(a, b int) int { return a * b },
	}

	for i := range input {
		if op, ok := opChars[input[i]]; ok {
			leftRes := diffWaysToCompute(input[:i])
			rightRes := diffWaysToCompute(input[i+1:])

			for _, leftV := range leftRes {
				for _, rightV := range rightRes {
					res = append(res, op(leftV, rightV))
				}
			}
		}
	}

	if len(res) == 0 {
		num, _ := strconv.Atoi(input)
		return []int{num}
	}

	return res
}
