/*
__author__ = 'lawtech'
__date__ = '2020/07/08 3:11 下午'
*/

package _670

import "strconv"

// 贪心
func maximumSwap(num int) int {
	hash := make(map[int]int)
	numStr := strconv.Itoa(num)
	chArr := []byte(numStr)
	for i := 0; i < len(numStr); i++ {
		hash[int(numStr[i]-'0')] = i
	}

	for i := 0; i < len(numStr); i++ {
		t := int(numStr[i] - '0')
		for j := 9; j > t; j-- {
			if hash[j] > i {
				chArr[i], chArr[hash[j]] = chArr[hash[j]], chArr[i]
				res, _ := strconv.Atoi(string(chArr))
				return res
			}
		}
	}

	return num
}
