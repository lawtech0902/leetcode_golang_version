/*
__author__ = 'lawtech'
__date__ = '2020/04/15 11:04 上午'
*/

package _274

import "sort"

// 升序排序，再倒序扫描
func hIndex(citations []int) int {
	sort.Ints(citations)
	i := 0
	for i < len(citations) && citations[len(citations)-1-i] > i {
		i++
	}

	return i
}
