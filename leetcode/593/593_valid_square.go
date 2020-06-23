/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-23 11:06:32
 */

package _593

import "sort"

// 排序
func validSquare(p1, p2, p3, p4 []int) bool {
	p := [][]int{p1, p2, p3, p4}
	sort.Slice(p, func(i, j int) bool {
		if p[i][1] < p[j][1] || p[i][1] == p[j][1] && p[i][0] <= p[j][0] {
			return true
		}

		return false
	})

	return dist(p[0], p[1]) != 0 && dist(p[0], p[1]) == dist(p[1], p[3]) && dist(p[1], p[3]) == dist(p[3], p[2]) && dist(p[3], p[2]) == dist(p[2], p[0]) && dist(p[0], p[3]) == dist(p[1], p[2])
}

func dist(p1, p2 []int) int {
	return (p2[1]-p1[1])*(p2[1]-p1[1]) + (p2[0]-p1[0])*(p2[0]-p1[0])
}
