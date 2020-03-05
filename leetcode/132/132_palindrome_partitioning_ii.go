/*
__author__ = 'lawtech'
__date__ = '2020/03/05 4:32 下午'
*/

package _132

import "math"

func minCut(s string) int {
	// p[i][j]表示从字符i到j是否为一个回文字符串。dp[i]表示从第i个字符到最后一个字符，最少的分割次数下，有多少个回文字符串，即分割次数+1。
	size := len(s)

	dp := make([]int, size+1)
	p := make([][]bool, size)
	for i := range p {
		p[i] = make([]bool, size)
	}

	for i := 0; i < size+1; i++ {
		dp[i] = size - i
	}

	for i := size - 1; i > -1; i-- {
		for j := i; j < size; j++ {
			if s[i] == s[j] && (((j - i) < 2) || p[i+1][j-1]) {
				p[i][j] = true
				dp[i] = int(math.Min(float64(1+dp[j+1]), float64(dp[i])))
			}
		}
	}

	return dp[0] - 1
}
