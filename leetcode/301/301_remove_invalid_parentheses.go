/*
__author__ = 'lawtech'
__date__ = '2020/04/20 3:38 下午'
*/

package _301

func removeInvalidParentheses(s string) []string {
	res := make([]string, 0)
	braces := []byte{'(', ')'}
	if len(s) == 0 {
		res = append(res, "")
		return res
	}

	visited := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, s)
	visited[s] = true

	found := false
	for len(queue) != 0 {
		s := queue[0]
		queue = queue[1:]

		if isValid(s, braces) {
			found = true
			res = append(res, s)
		}

		if found {
			continue
		}

		for i := 0; i < len(s); i++ {
			if s[i] != braces[0] && s[i] != braces[1] {
				continue
			}

			t := s[0:i] + s[i+1:]
			if _, ok := visited[t]; !ok {
				queue = append(queue, t)
				visited[t] = true
			}
		}
	}

	return res
}

func isValid(s string, braces []byte) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != braces[0] && s[i] != braces[1] {
			continue
		} else if s[i] == braces[0] {
			count++
		} else if s[i] == braces[1] {
			count--
		}

		if count < 0 {
			return false
		}
	}

	return count == 0
}
