/*
__author__ = 'lawtech'
__date__ = '2020/06/15 1:30 下午'
*/

package _541

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	i := 0
	for i < len(bytes) {
		l, r := i, i+k-1
		if r >= len(bytes) {
			r = len(bytes) - 1
		}

		for l < r {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
		}

		i += 2 * k
	}

	return string(bytes)
}
