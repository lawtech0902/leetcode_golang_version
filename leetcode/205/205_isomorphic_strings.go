/*
__author__ = 'lawtech'
__date__ = '2020/03/23 2:37 下午'
*/

package _205

import "strings"

// 遍历字符串
func isIsomorphic(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s[i+1:], s[i]) != strings.IndexByte(t[i+1:], t[i]) {
			return false
		}
	}

	return true
}
