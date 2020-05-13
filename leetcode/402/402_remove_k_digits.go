/*
__author__ = 'lawtech'
__date__ = '2020/05/13 10:16 上午'
*/

package _402

import "strings"

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		c := num[i]
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c { // 维持单调非递减
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, c)
	}

	stack = stack[:len(stack)-k] // 删除末尾数字
	s := string(stack)
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}

	return s
}
