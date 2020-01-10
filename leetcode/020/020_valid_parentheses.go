/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-01-03 16:19:42
 */

package _20

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
