/*
__author__ = 'lawtech'
__date__ = '2020/05/12 10:59 上午'
*/

package _390

// 恶心的数学规律题
func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}

	return 2 * (n/2 + 1 - lastRemaining(n/2))
}
