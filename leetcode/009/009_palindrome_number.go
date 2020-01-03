package main

import (
	"strconv"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 上午12:40'
*/

func isPalindrome(x int) interface{} {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
