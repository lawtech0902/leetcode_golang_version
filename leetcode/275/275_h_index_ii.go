/*
__author__ = 'lawtech'
__date__ = '2020/04/15 11:24 上午'
*/

package _275

func hIndex(citations []int) int {
	i := 0
	for i < len(citations) && citations[len(citations)-1-i] > i {
		i++
	}

	return i
}
