/*
__author__ = 'lawtech'
__date__ = '2020/07/23 2:19 下午'
*/

package _779

import "math/bits"

// 递归，父子递推，第 K 个数字是上一行第 (K+1) / 2 个数字生成的。
// 如果上一行的数字为 0，被生成的数字为 1 - (K%2)，如果上一行的数字为 1，被生成的数字为 K%2。
func kthGrammar(N, K int) int {
	if N == 1 {
		return 0
	}

	return (^K & 1) ^ kthGrammar(N-1, (K+1)/2)
}

// binary count
func kthGrammar1(N, K int) int {
	return bits.OnesCount(uint(K-1)) % 2
}
