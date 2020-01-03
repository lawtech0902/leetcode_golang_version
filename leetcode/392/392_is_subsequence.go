package _392

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:09'
*/

func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
