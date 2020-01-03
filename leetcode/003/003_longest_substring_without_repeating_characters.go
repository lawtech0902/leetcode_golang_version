package _03

/*
__author__ = 'lawtech'
__date__ = '2018/8/17 下午11:37'
*/

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}
