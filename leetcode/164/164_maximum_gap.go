/*
__author__ = 'lawtech'
__date__ = '2020/03/13 4:18 下午'
*/

package _164

import (
	"math"
	"sort"
)

func maximumGap(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		res = int(math.Max(float64(res), float64(nums[i+1]-nums[i])))
	}

	return res
}
