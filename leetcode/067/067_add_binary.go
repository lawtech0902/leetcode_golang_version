/*
__author__ = 'lawtech'
__date__ = '2020/1/16 3:40 下午'
*/

package _67

import "strconv"

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	res := ""
	flag := 0 // 进位标识
	cur := 0

	for i >= 0 || j >= 0 {
		intA, intB := 0, 0
		if i >= 0 {
			intA = int(a[i] - '0')
		}
		if j >= 0 {
			intB = int(b[j] - '0')
		}

		cur = intA + intB + flag
		flag = 0
		if cur >= 2 {
			flag = 1
			cur = cur - 2
		}

		curStr := strconv.Itoa(cur)
		res = curStr + res
		i--
		j--
	}

	if flag == 1 {
		res = "1" + res
	}

	return res
}
