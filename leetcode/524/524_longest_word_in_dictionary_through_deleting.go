/*
__author__ = 'lawtech'
__date__ = '2020/06/08 10:24 上午'
*/

package _524

func findLongestWord(s string, d []string) string {
	var res string
	for _, v := range d {
		if isSubsequence(v, s) {
			if len(v) > len(res) || len(v) == len(res) && v < res {
				res = v
			}
		}
	}

	return res
}

func isSubsequence(p, s string) bool {
	i := 0
	for j := 0; j < len(s) && i < len(p); j++ {
		if p[i] == s[j] {
			i++
		}
	}

	return i == len(p)
}
