/*
__author__ = 'lawtech'
__date__ = '2020/05/20 3:52 下午'
*/

package _443

import "fmt"

// 双指针读写
func compress(chars []byte) int {
	scan, write := 0, 0
	l := len(chars)

	for scan < l {
		count := 0
		chars[write] = chars[scan]
		for scan < l && chars[write] == chars[scan] {
			count++
			scan++
		}

		if count > 1 {
			tmp := fmt.Sprintf("%d", count)
			for _, c := range []byte(tmp) {
				write++
				chars[write] = c
			}
		}

		write++
	}

	return write
}
