/*
__author__ = 'lawtech'
__date__ = '2020/05/28 1:33 下午'
*/

package _481

import "strings"

func magicalString(n int) int {
	s := "122"
	// pre指向需要构造的次数，now指向应该构造的字母
	pre, now := 2, 2

	for {
		if now >= n {
			return strings.Count(s[:n], "1")
		}

		if s[pre] == '2' {
			if s[now] == '2' {
				s += strings.Repeat("1", 2)
			} else {
				s += strings.Repeat("2", 2)
			}

			now += 2
		} else {
			if s[now] == '1' {
				s += "2"
			} else {
				s += "1"
			}

			now += 1
		}

		pre += 1
	}
}
