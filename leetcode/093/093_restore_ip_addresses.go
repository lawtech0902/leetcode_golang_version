/*
__author__ = 'lawtech'
__date__ = '2020/1/19 1:56 下午'
*/

package _93

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	size := len(s)
	var res []string

	if size < 4 || size > 12 {
		return res
	}

	for i := 1; i < size-2; i++ {
		first, _ := strconv.Atoi(s[:i])
		if first > 255 {
			break
		}

		for j := i + 1; j < size-1; j++ {
			second, _ := strconv.Atoi(s[i:j])
			if second > 255 {
				break
			}

			for k := j + 1; k < size; k++ {
				third, _ := strconv.Atoi(s[j:k])
				if third > 255 {
					break
				}

				last, _ := strconv.Atoi(s[k:])
				if last > 255 {
					continue
				}

				ip := strings.Join([]string{strconv.Itoa(first), strconv.Itoa(second), strconv.Itoa(third), strconv.Itoa(last)}, ".")
				if len(ip) == size+3 {
					res = append(res, ip)
				}
			}
		}
	}

	return res
}
