/*
__author__ = 'lawtech'
__date__ = '2020/06/03 2:15 下午'
*/

package _506

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	var res []string
	temp := make([]int, len(nums))
	copy(temp, nums)
	hash := make(map[int]string)

	sort.Sort(sort.Reverse(sort.IntSlice(temp)))
	for i, v := range temp {
		switch i {
		case 0:
			hash[v] = "Gold Medal"
		case 1:
			hash[v] = "Silver Medal"
		case 2:
			hash[v] = "Bronze Medal"
		default:
			hash[v] = strconv.Itoa(i + 1)
		}
	}

	for _, num := range nums {
		res = append(res, hash[num])
	}

	return res
}
