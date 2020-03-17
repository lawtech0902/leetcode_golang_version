/*
__author__ = 'lawtech'
__date__ = '2020/03/17 2:04 下午'
*/

package _187

func findRepeatedDnaSequences(s string) []string {
	var res []string
	m := make(map[string]int)

	for i := 0; i < len(s)-9; i++ {
		subStr := s[i : i+10]
		if _, exists := m[subStr]; exists {
			m[subStr] += 1
		} else {
			m[subStr] = 1
		}
	}

	for subStr, cnt := range m {
		if cnt > 1 {
			res = append(res, subStr)
		}
	}

	return res
}
