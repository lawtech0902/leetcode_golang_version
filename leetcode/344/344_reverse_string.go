/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午10:00'
*/

package _344

// 双指针
func reverseString(s []byte) string {
	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return string(s)
}
