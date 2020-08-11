/*
__author__ = 'lawtech'
__date__ = '2020/08/11 10:00 上午'
*/

package _63

import "math"

func maxProfit(prices []int) int {
	cost, profit := math.MaxInt32, 0
	for _, price := range prices {
		cost = min(cost, price)
		profit = max(profit, price-cost)
	}

	return profit
}

func min(values ...int) int {
	minValue := math.MaxInt32
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}

func max(values ...int) int {
	maxValue := math.MinInt32
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}
