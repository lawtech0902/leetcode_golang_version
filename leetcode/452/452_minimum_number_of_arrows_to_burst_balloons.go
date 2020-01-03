package _452

import "sort"

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午1:53'
*/

type balloons [][]int

func (b balloons) Len() int {
	return len(b)
}

// balloons 实现了 sort.Interface 接口
// 以end的大小来排序
func (b balloons) Less(i, j int) bool {
	return b[i][1] < b[j][1]
}

func (b balloons) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func findMinArrowShots(points [][]int) int {
	size := len(points)

	if size == 0 {
		return 0
	}

	sort.Sort(balloons(points))

	res := 0
	end := points[0][1]

	for i := 1; i < size; i++ {
		if points[i][0] <= end {
			continue
		}
		res++
		end = points[i][1]
	}

	res++

	return res
}
