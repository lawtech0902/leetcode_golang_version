/*
__author__ = 'lawtech'
__date__ = '2020/04/23 10:19 上午'
*/

package _319

func bulbSwitch(n int) int {
	i := 1
	for i*i <= n {
		i++
	}

	return i - 1
}
