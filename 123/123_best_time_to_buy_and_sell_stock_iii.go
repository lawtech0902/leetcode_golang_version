package _123

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 下午10:59'
*/

func maxProfit(prices []int) int {
	firstBuy, firstSell := math.MinInt32, 0
	secondBuy, secondSell := math.MinInt32, 0

	for _, curPrice := range prices {
		if firstBuy < -curPrice {
			firstBuy = -curPrice
		}
		if firstSell < firstBuy+curPrice {
			firstSell = firstBuy + curPrice
		}
		if secondBuy < firstSell-curPrice {
			secondBuy = firstSell - curPrice
		}
		if secondSell < secondBuy+curPrice {
			secondSell = secondBuy + curPrice
		}
	}

	return secondSell
}
