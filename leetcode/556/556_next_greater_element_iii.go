/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-16 10:55:36
 */

package _556

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 参考031_next_permutation
func nextGreaterElement(n int) int {
	str := strconv.Itoa(n)
	if len(str) == 1 {
		return -1
	}

	strSlice := make([]string, 0)
	for i := 0; i < len(str); i++ {
		strSlice = append(strSlice, string(str[i]))
	}

	left := len(strSlice) - 2
	for left >= 0 && strSlice[left] >= strSlice[left+1] {
		left--
	}

	if left == -1 {
		return -1
	}

	sort.Strings(strSlice[left+1:])

	right := search(strSlice, left+1, strSlice[left])
	strSlice[left], strSlice[right] = strSlice[right], strSlice[left]

	res, _ := strconv.Atoi(strings.Join(strSlice, ""))
	if res == n || res > math.MaxInt32 {
		return -1
	}

	return res
}

func search(a []string, l int, target string) int {
	r := len(a) - 1
	l--

	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}
