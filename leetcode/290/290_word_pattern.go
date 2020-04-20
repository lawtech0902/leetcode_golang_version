/*
__author__ = 'lawtech'
__date__ = '2020/04/20 1:28 下午'
*/

package _290

import "strings"

func wordPattern(pattern string, str string) bool {
	strArr := strings.Split(str, " ")
	if len(pattern) != len(strArr) {
		return false
	}

	patternMap := make(map[byte]string)
	mappedWord := make(map[string]bool)

	for i := 0; i < len(pattern); i++ {
		matchWord, ok := patternMap[pattern[i]]
		if !ok {
			if mapped := mappedWord[strArr[i]]; mapped {
				return false
			}

			patternMap[pattern[i]] = strArr[i]
			mappedWord[strArr[i]] = true
			continue
		}

		if strArr[i] != matchWord {
			return false
		}
	}

	return true
}
