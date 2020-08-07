/*
__author__ = 'lawtech'
__date__ = '2020/08/07 2:34 下午'
*/

package _58

// 切片
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
