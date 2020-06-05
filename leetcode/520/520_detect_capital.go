/*
__author__ = 'lawtech'
__date__ = '2020/06/05 11:22 上午'
*/

package _520

import "strings"

func detectCapitalUse(word string) bool {
	return strings.ToLower(word) == word || strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}
