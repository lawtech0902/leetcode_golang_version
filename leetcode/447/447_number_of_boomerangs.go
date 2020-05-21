/*
__author__ = 'lawtech'
__date__ = '2020/05/21 10:14 上午'
*/

package _447

func numberOfBoomerangs(points [][]int) int {
	count := 0

	for i := 0; i < len(points); i++ {
		distances := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			tmp := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) +
				(points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			distances[tmp]++
		}

		for _, v := range distances {
			count += v * (v - 1)
		}
	}

	return count
}
