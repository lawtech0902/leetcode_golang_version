/*
__author__ = 'lawtech'
__date__ = '2020/07/16 10:55 上午'
*/

package _718

// dp
func findLength(A, B []int) int {
	l1, l2 := len(A), len(B)
	res := 0
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				res = max(res, dp[i+1][j+1])
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
