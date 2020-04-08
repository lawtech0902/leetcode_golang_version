/*
__author__ = 'lawtech'
__date__ = '2020/04/08 1:50 下午'
*/

package _231

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
