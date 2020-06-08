/*
__author__ = 'lawtech'
__date__ = '2020/06/08 10:59 上午'
*/

package _526

var res int
var hmap [][]int

func countArrangement(N int) int {
	res = 0
	hmap = make([][]int, N+1)
	for i := 1; i <= N; i++ {
		hmap[i] = make([]int, 0)
		for j := 1; j <= N; j++ {
			if j%i == 0 || i%j == 0 {
				hmap[i] = append(hmap[i], j)
			}
		}
	}

	help(N, 0, N)
	return res
}

func help(i int, visited int, N int) {
	if i == 0 {
		res++
		return
	}
	for _, e := range hmap[i] {
		if visited&(1<<e) > 0 {
			continue
		}
		help(i-1, visited|(1<<e), N)
	}
}
