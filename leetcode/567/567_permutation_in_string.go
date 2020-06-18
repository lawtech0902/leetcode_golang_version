/*
__author__ = 'lawtech'
__date__ = '2020/06/18 10:02 上午'
*/

package _567

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	hash1 := make([]int, 26)
	hash2 := make([]int, 26)
	for i := range s1 {
		hash1[s1[i]-'a']++
		hash2[s2[i]-'a']++
	}

	if match(hash1, hash2) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		hash2[s2[i]-'a']++
		hash2[s2[i-len(s1)]-'a']--
		if s2[i] != s2[i-len(s1)] && match(hash1, hash2) {
			return true
		}
	}

	return false
}

func match(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
