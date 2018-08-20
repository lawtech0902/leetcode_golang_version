package _122

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:30'
*/

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
