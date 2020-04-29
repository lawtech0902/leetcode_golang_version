/*
__author__ = 'lawtech'
__date__ = '2020/04/29 9:39 上午'
*/

package _342

// 若 x 为 2 的幂且 x%3 == 1，则 x 为 4 的幂。
func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num%3 == 1
}
