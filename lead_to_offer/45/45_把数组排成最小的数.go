/*
__author__ = 'lawtech'
__date__ = '2020/08/04 4:39 下午'
*/

package _45

import (
	"fmt"
	"sort"
	"strings"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compareNum(nums[i], nums[j])
	})
	var res strings.Builder
	for _, num := range nums {
		res.WriteString(fmt.Sprintf("%d", num))
	}

	return res.String()
}

func compareNum(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	return str1 < str2
}
