/*
__author__ = 'lawtech'
__date__ = '2020/05/12 10:13 上午'
*/

package _387

import "math"

// arr
func firstUniqChar(s string) int {
	cache := [26]int{}
	for _, ch := range s {
		cache[ch-'a']++
	}

	for i := range s {
		if cache[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// hashmap
func firstUniqChar1(s string) int {
	cache := make(map[int32]int)
	for i, ch := range s {
		if _, ok := cache[ch]; ok {
			cache[ch] = -1
			continue
		}

		cache[ch] = i
	}

	first := math.MaxInt32
	for _, v := range cache {
		if v != -1 && first > v {
			first = v
		}
	}

	if first == math.MaxInt32 {
		return -1
	}

	return first
}
