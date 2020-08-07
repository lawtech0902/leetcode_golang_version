/*
__author__ = 'lawtech'
__date__ = '2020/08/07 2:22 下午'
*/

package _58

import "strings"

func reverseWords(s string) string {
	strArr := strings.Fields(s)
	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}

	return strings.Join(strArr, " ")
}
