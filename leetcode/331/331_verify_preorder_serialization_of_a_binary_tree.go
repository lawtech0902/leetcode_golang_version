/*
__author__ = 'lawtech'
__date__ = '2020/04/26 1:34 下午'
*/

package _331

import "strings"

// O(n) O(n)
func isValidSerialization(preorder string) bool {
	slots := 1
	for _, node := range strings.Split(preorder, ",") {
		slots--
		if slots < 0 {
			return false
		}

		if node != "#" {
			slots += 2
		}
	}

	return slots == 0
}

// O(n) O(1)
func isValidSerialization1(preorder string) bool {
	slots := 1
	var prev rune

	for _, ch := range preorder {
		if ch == ',' {
			slots--
			if slots < 0 {
				return false
			}

			if prev != '#' {
				slots += 2
			}
		}

		prev = ch
	}

	// last node
	if prev != '#' {
		slots++
	} else {
		slots--
	}

	return slots == 0
}
