/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:23'
*/

package _121

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	low := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}

		maxProfit = int(math.Max(float64(maxProfit), float64(prices[i]-low)))
	}

	return maxProfit
}
