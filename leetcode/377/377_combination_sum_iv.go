/*
__author__ = 'lawtech'
__date__ = '2018/8/21 下午8:28'
*/

package _377

// dp[x]表示生成数字x的所有可能的组合方式的个数。
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i-v >= 0 {
				dp[i] += dp[i-v]
			}
		}
	}

	return dp[target]
}
