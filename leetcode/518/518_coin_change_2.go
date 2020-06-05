/*
__author__ = 'lawtech'
__date__ = '2020/06/05 10:34 上午'
*/

package _518

// dp
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] += dp[x-coin]
		}
	}

	return dp[amount]
}
