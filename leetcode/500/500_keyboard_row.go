/*
__author__ = 'lawtech'
__date__ = '2020/06/03 10:38 上午'
*/

package _500

import "strings"

func findWords(words []string) []string {
	indexes := [26]int{}
	for _, v := range "qwertyuiop" {
		indexes[v-'a'] = 1
	}

	for _, v := range "asdfghjkl" {
		indexes[v-'a'] = 2
	}

	for _, v := range "zxcvbnm" {
		indexes[v-'a'] = 3
	}

	res := make([]string, 0)
	for _, word := range words {
		temp := strings.ToLower(word)
		flag := true
		for _, c := range temp {
			if indexes[c-'a'] != indexes[temp[0]-'a'] {
				flag = false
				break
			}
		}

		if flag {
			res = append(res, word)
		}
	}

	return res
}
