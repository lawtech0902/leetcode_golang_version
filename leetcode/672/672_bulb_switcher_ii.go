/*
__author__ = 'lawtech'
__date__ = '2020/07/08 4:27 下午'
*/

package _672

func flipLights(n, m int) int {
	n = min(n, 3)
	switch m {
	case 0:
		return 1
	case 1:
		return []int{2, 3, 4}[n-1]
	case 2:
		return []int{2, 4, 7}[n-1]
	default:
		return []int{2, 4, 8}[n-1]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
