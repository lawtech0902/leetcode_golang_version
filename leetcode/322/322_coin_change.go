package _322

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 上午11:31'
*/

/*
if dp[x+c] < 0 || dp[x+c] > dp[x]+1 {
	dp[x+c] = dp[x] + 1
}
其中dp[x]代表凑齐金额x所需的最少硬币数，c为可用的硬币面值

初始令dp[0] = 0
*/

func coinChange(coins []int, amount int) int {
	// 傻逼一样的数组声明
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}

	for x := 0; x < amount; x++ {
		if dp[x] < 0 {
			continue
		}
		for _, c := range coins {
			if x+c > amount {
				continue
			}

			if dp[x+c] < 0 || dp[x+c] > dp[x]+1 {
				dp[x+c] = dp[x] + 1
			}
		}
	}

	return dp[amount]
}
