/*
__author__ = 'lawtech'
__date__ = '2020/03/11 10:18 上午'
*/

package _149

import "math"

type Point struct {
	x, y int
}

func maxPoints(points [][]int) int {
	n := len(points)
	if n < 2 {
		return n
	}

	counts := map[Point]int{} // 统计点的重复个数
	for _, v := range points {
		point := Point{v[0], v[1]}
		if _, ok := counts[point]; ok {
			counts[point]++
		} else {
			counts[point] = 1
		}
	}

	points = [][]int{} // 只保留不重复的点
	for p := range counts {
		points = append(points, []int{p.x, p.y})
	}

	// 经过每个点的直线的斜率，以及该直线上的点的个数
	lines := make([]map[float64]int, n)
	for i, _ := range lines {
		lines[i] = make(map[float64]int)
	}

	res := 0
	for i, v := range points {
		countI := counts[Point{v[0], v[1]}]
		res = int(math.Max(float64(res), float64(countI))) // 特殊情况：所有点全部相同

		for j := 0; j < i; j++ {
			countJ := counts[Point{points[j][0], points[j][1]}]

			// 计算斜率
			k := 0.0
			deltaX := points[i][0] - points[j][0]
			deltaY := points[i][1] - points[j][1]

			if deltaX == 0 {
				k = math.Inf(1)
			} else {
				k = float64(deltaY) / float64(deltaX)
			}

			if _, exist := lines[j][k]; exist {
				lines[i][k] = lines[j][k] + countI
			} else {
				lines[i][k] = countI + countJ
			}

			res = int(math.Max(float64(res), float64(lines[i][k])))
		}
	}

	return res
}
