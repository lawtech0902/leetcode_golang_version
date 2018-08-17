package _05

/*
__author__ = 'lawtech'
__date__ = '2018/8/17 下午11:44'
*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 最长回文的首字符索引，长度
	begin, maxLen := 0, 1

	// 中心扩展法
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}

		// b 代表回文的首字符索引号，
		// e 代表回文的尾字符索引号，
		// i 代表回文`正中间段`首字符的索引号
		b, e := i, i

		// 中心相同字符扩展
		for e < len(s)-1 && s[e+1] == s[e] {
			e ++
		}

		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e ++
			b --
		}

		newLen := e + 1 - b
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin:begin+maxLen]
}
