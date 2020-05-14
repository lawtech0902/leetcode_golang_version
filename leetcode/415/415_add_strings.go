/*
__author__ = 'lawtech'
__date__ = '2020/05/14 10:29 上午'
*/

package _415

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	builder := strings.Builder{}
	i, j, carry := len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
		} else {
			n1 = 0
		}

		if j >= 0 {
			n2 = int(num2[j] - '0')
		} else {
			n2 = 0
		}

		tmp := n1 + n2 + carry
		carry = tmp / 10
		builder.WriteString(strconv.Itoa(tmp % 10))
		i, j = i-1, j-1
	}

	if carry == 1 {
		builder.WriteString(strconv.Itoa(1))
	}

	return reverseString(builder.String())
}

func reverseString(s string) string {
	bytes := []byte(s)
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}

	return string(bytes)
}
