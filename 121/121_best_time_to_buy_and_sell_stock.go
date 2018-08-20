package _121

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:23'
*/

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	low := prices[0]
	max_profit := 0

	for price := range prices {
		if price < low {
			low = price
		}

		max_profit = int(math.Max(float64(price-low), float64(max_profit)))
	}

	return max_profit
}
