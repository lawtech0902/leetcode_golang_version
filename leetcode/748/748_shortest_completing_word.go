/*
__author__ = 'lawtech'
__date__ = '2020/07/21 2:22 下午'
*/

package _748

func shortestCompletingWord(licensePlate string, words []string) string {
	var res string
	src := count(licensePlate)
	for _, word := range words {
		if !isValid(src, word) {
			continue
		}

		if res == "" || len(word) < len(res) {
			res = word
		}
	}

	return res
}

func count(word string) []int {
	res := make([]int, 26)
	for _, c := range word {
		if 'A' <= c && c <= 'Z' {
			res[c-'A']++
		} else if 'a' <= c && c <= 'z' {
			res[c-'a']++
		}
	}

	return res
}

func isValid(src []int, word string) bool {
	dest := count(word)
	for i := range src {
		if src[i] > dest[i] {
			return false
		}
	}

	return true
}
