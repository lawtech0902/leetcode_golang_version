/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午8:57'
*/

package _28

func strStr(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)

	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}

	return -1
}
