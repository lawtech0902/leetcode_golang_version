/*
__author__ = 'lawtech'
__date__ = '2020/05/12 1:32 下午'
*/

package _394

import "strings"

// 双辅助栈
func decodeString(s string) string {
	stackNums := make([]int, 0)
	stackStr := make([]string, 0)
	var (
		res string
		num int
	)

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9': // num
			num = 10*num + int(s[i]) - '0'
		case s[i] == '[':
			stackNums = append(stackNums, num)
			num = 0
			stackStr = append(stackStr, res)
			res = ""
		case s[i] == ']':
			tmp := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
			count := stackNums[len(stackNums)-1]
			stackNums = stackNums[:len(stackNums)-1]
			res = tmp + strings.Repeat(res, count)
		default:
			res += string(s[i])
		}
	}

	return res
}
