/*
__author__ = 'lawtech'
__date__ = '2020/07/16 10:22 上午'
*/

package _714

// dp
func maxProfit(prices []int, fee int) int {
	notHold, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		notHold = max(notHold, hold+prices[i]-fee)
		hold = max(hold, notHold-prices[i])
	}

	return notHold
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
