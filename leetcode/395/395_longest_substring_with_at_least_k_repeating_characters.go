/*
__author__ = 'lawtech'
__date__ = '2020/05/12 2:12 下午'
*/

package _395

// 多路分治
func longestSubstring(s string, k int) int {
	counts, length := make(map[byte]int), len(s)
	if length == 0 {
		return 0
	}

	for i := range s {
		counts[s[i]]++
	}

	max := 0
	for i := 0; i < length; {
		for i < length && counts[s[i]] < k {
			i++
		}

		start := i
		for i < length && counts[s[i]] >= k {
			i++
		}

		end := i
		if start == 0 && end == length {
			return length
		}

		res := longestSubstring(s[start:end], k)
		if max < res {
			max = res
		}
	}

	return max
}
