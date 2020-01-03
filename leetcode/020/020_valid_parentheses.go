package _20

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 上午1:30'
*/

func isValid(s string) bool {
	var stack []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			l := len(stack)
			if l == 0 {
				return false
			}

			v := stack[l-1]
			stack = stack[:l-1]

			if c == ')' && v != '(' {
				return false
			}

			if c == ']' && v != '[' {
				return false
			}

			if c == '}' && v != '{' {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
