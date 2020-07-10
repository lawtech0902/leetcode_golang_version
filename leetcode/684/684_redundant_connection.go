/*
__author__ = 'lawtech'
__date__ = '2020/07/10 10:43 上午'
*/

package _684

// 一维数组p[i] = j表示第i个节点的前驱节点为j，连通树中根节点k满足p[k] = k
func findRedundantConnection(edges [][]int) []int {
	res := make([]int, 2)
	p := make([]int, len(edges)+1)
	for i := 1; i <= len(edges); i++ {
		p[i] = i
	}

	for i := 0; i < len(edges); i++ {
		a := find(edges[i][0], p)
		b := find(edges[i][1], p)
		if a == b {
			res = []int{edges[i][0], edges[i][1]}
		} else {
			// 合并
			union(a, b, p)
		}
	}

	return res
}

// 包含树压缩动作，将每个子集合树变为二级
func find(x int, p []int) int {
	if p[x] != x {
		p[x] = find(p[x], p)
	}

	return p[x]
}

// 对两个集合按树根节点（子集的领导节点）合并，x, y为两个集合的任意两个节点
func union(x, y int, p []int) {
	xRoot := find(x, p)
	yRoot := find(y, p)
	if xRoot != yRoot {
		p[xRoot] = yRoot
	}
}
