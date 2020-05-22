/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午1:53'
*/

package _452

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1
	position := points[0][1]
	for i := 0; i < len(points); i++ {
		if points[i][0] <= position {
			continue
		}

		res++
		position = points[i][1]
	}

	return res
}
