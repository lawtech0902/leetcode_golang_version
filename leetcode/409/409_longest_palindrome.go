/*
__author__ = 'lawtech'
__date__ = '2020/05/13 1:54 下午'
*/

package _409

func longestPalindrome(s string) int {
	var count [123]int
	for i := range s {
		count[s[i]]++
	}

	res, odd := 0, 0
	for _, v := range count {
		res += v
		if v%2 == 1 {
			res -= 1
			odd += 1
		}
	}

	if odd != 0 {
		return res + 1
	}

	return res
}
