/*
__author__ = 'lawtech'
__date__ = '2020/03/13 4:35 下午'
*/

package _165

import (
	"math"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1Arr := strings.Split(version1, ".")
	v2Arr := strings.Split(version2, ".")
	size1 := len(v1Arr)
	size2 := len(v2Arr)
	maxLen := int(math.Max(float64(size1), float64(size2)))

	for i := 0; i < maxLen; i++ {
		v1Token := 0
		if i < size1 {
			v1Token, _ = strconv.Atoi(v1Arr[i])
		}

		v2Token := 0
		if i < size2 {
			v2Token, _ = strconv.Atoi(v2Arr[i])
		}

		if v1Token < v2Token {
			return -1
		}

		if v1Token > v2Token {
			return 1
		}
	}

	return 0
}
