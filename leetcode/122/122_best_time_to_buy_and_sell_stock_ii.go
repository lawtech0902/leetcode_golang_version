/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:30'
*/

package _122

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
