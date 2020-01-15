/*
__author__ = 'lawtech'
__date__ = '2020/1/15 3:51 下午'
*/

package _49

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}

	for _, word := range strs {
		wordSlice := strings.Split(word, "")
		sort.Strings(wordSlice)
		sortedWord := strings.Join(wordSlice[:], "")
		if _, ok := dict[sortedWord]; !ok {
			dict[sortedWord] = []string{word}
		} else {
			dict[sortedWord] = append(dict[sortedWord], word)
		}
	}

	var res [][]string
	for _, v := range dict {
		res = append(res, v)
	}

	return res
}
