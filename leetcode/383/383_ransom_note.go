/*
__author__ = 'lawtech'
__date__ = '2020/05/09 2:29 下午'
*/

package _383

func canConstruct(ransomNote string, magazine string) bool {
	data := make([]int, 26)
	for _, c := range magazine {
		data[c-'a']++
	}

	for _, c := range ransomNote {
		if data[c-'a'] == 0 {
			return false
		}

		data[c-'a']--
	}

	return true
}
