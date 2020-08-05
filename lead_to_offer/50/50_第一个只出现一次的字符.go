/*
__author__ = 'lawtech'
__date__ = '2020/08/05 2:12 下午'
*/

package _50

// map
func firstUniqChar(s string) byte {
	hash := make(map[byte]int)
	for i := range s {
		hash[s[i]]++
	}

	for i := range s {
		if hash[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}

// slice
func firstUniqChar1(s string) byte {
	var list [26]int
	for i := range s {
		list[s[i]-'a']++
	}

	for i := range s {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}
