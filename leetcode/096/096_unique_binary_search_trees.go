/*
__author__ = 'lawtech'
__date__ = '2020/2/4 2:46 下午'
*/

package _96

func numTrees(n int) int {
	// 卡特兰数，这种题目记记答案，我等泛泛之辈实在无能为力
	dp := []int{1, 1, 2}
	if n < 2 {
		return dp[n]
	}

	dp = append(dp, make([]int, n-2)...)

	for i := 3; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
