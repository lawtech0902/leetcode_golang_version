/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-06 09:57:02
 */

package _647

func countSubstrings(s string) int {
	var (
		res   int
		count func(s string, start, end int)
	)

	count = func(s string, start, end int) {
		for start >= 0 && end < len(s) && s[start] == s[end] {
			res++
			start--
			end++
		}
	}

	for i := 0; i < len(s); i++ {
		count(s, i, i)   // 回文串长度为奇数
		count(s, i, i+1) //回文串长度为偶数
	}

	return res
}
