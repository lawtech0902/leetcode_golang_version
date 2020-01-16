/*
__author__ = 'lawtech'
__date__ = '2020/1/16 11:10 上午'
*/

package _58

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	sArray := strings.Split(s, " ")
	return len(sArray[len(sArray)-1])
}
