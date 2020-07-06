/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-06 10:18:01
 */

package _648

import (
	"fmt"
	"strings"
)

// 前缀哈希
func replaceWords(dict []string, sentence string) string {
	hash := make(map[string]interface{})
	max, min := 0, 1000
	for _, v := range dict {
		hash[v] = nil
		l := len(v)
		if l > max {
			max = l
		}

		if l < min {
			min = l
		}
	}

	fields := strings.Fields(sentence)
	for i, v := range fields {
		for j := min; j < len(v) && j <= max; j++ {
			vs := v[:j]
			if _, exist := hash[vs]; exist {
				fmt.Println(fields[i], vs)
				fields[i] = vs
				break
			}
		}
	}

	return strings.Join(fields, " ")
}
