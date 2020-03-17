/*
__author__ = 'lawtech'
__date__ = '2020/03/17 9:40 上午'
*/

package _171

func titleToNumber(s string) int {
	// 26进制转10进制
	sum, base := 0, 1
	for i := len(s) - 1; i > -1; i-- {
		sum += int(s[i]-'A'+1) * base
		base *= 26
	}

	return sum
}
