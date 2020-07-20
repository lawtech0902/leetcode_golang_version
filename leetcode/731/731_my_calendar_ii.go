/*
__author__ = 'lawtech'
__date__ = '2020/07/20 10:36 上午'
*/

package _731

// 暴力
type MyCalendarTwo struct {
	calendar [][]int
	overlaps [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		calendar: make([][]int, 0),
		overlaps: make([][]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, v := range this.overlaps {
		if start < v[1] && end > v[0] {
			return false
		}
	}

	for _, v := range this.calendar {
		if start < v[1] && end > v[0] {
			this.overlaps = append(this.overlaps, []int{max(start, v[0]), min(end, v[1])})
		}
	}

	this.calendar = append(this.calendar, []int{start, end})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
