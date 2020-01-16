/*
__author__ = 'lawtech'
__date__ = '2020/1/16 4:23 下午'
*/

package _71

func simplifyPath(path string) string {
	var stack []string
	i := 0
	res := ""

	for i < len(path) {
		end := i + 1
		for end < len(path) && path[end] != '/' {
			end += 1
		}

		sub := path[i+1 : end]
		if len(sub) > 0 {
			if sub == ".." {
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				}
			} else if sub != "." {
				stack = append(stack, sub)
			}
		}

		i = end
	}

	if len(stack) == 0 {
		return "/"
	}

	for _, val := range stack {
		res += "/" + val
	}

	return res
}
