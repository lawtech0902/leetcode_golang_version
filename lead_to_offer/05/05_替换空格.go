/*
__author__ = 'lawtech'
__date__ = '2020/07/27 1:32 下午'
*/

package _5

import "strings"

// 内置函数
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// 遍历
func replaceSpace1(s string) string {
	var res strings.Builder
	for i := range s {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}

	return res.String()
}
