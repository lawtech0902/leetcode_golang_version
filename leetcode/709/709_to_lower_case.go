/*
__author__ = 'lawtech'
__date__ = '2020/07/15 1:58 下午'
*/

package _709

func toLowerCase(str string) string {
	r := []rune(str)
	diff := 'a' - 'A'
	for i := 0; i < len(r); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] += diff
		}
	}

	return string(r)
}
