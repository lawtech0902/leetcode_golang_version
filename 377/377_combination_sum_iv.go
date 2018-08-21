package _377

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 下午8:28'
*/

// dp[x]表示生成数字x的所有可能的组合方式的个数。
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i+num <= target {
				dp[i+num] += dp[i]
			}
		}
	}

	return dp[target]
}
