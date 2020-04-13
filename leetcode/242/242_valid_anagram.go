/*
__author__ = 'lawtech'
__date__ = '2020/04/13 3:37 下午'
*/

package _242

// HashMap O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
