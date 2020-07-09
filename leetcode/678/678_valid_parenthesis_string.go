/*
__author__ = 'lawtech'
__date__ = '2020/07/09 2:21 下午'
*/

package _678

// stack
func checkValidString(s string) bool {
	starStack, openStack := make([]int, 0), make([]int, 0)
	for i := range s {
		if s[i] == '*' {
			starStack = append(starStack, i)
		} else if s[i] == '(' {
			openStack = append(openStack, i)
		} else {
			if len(openStack) == 0 && len(starStack) == 0 {
				return false
			} else if len(openStack) > 0 {
				openStack = openStack[:len(openStack)-1]
			} else {
				starStack = starStack[:len(starStack)-1]
			}
		}
	}

	for len(openStack) > 0 {
		if len(starStack) == 0 || starStack[len(starStack)-1] < openStack[len(openStack)-1] {
			return false
		}

		starStack = starStack[:len(starStack)-1]
		openStack = openStack[:len(openStack)-1]
	}

	return true
}
