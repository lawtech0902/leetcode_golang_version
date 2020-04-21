/*
__author__ = 'lawtech'
__date__ = '2020/04/21 2:13 下午'
*/

package _309

import "math"

/*
sold：
前一天hold，当日卖出股票

hold： 可由两个情况转换来
前一天hold，当日rest
前一天rest，当日买入股票变为hold

rest：
前一天sold，当日必须rest
前一天rest，当日继续rest
*/
func maxProfit(prices []int) int {
	sold, rest, hold := 0, 0, math.MinInt32
	for _, p := range prices {
		preSold := sold
		sold = hold + p
		hold = max(hold, rest-p)
		rest = max(rest, preSold)
	}

	return max(sold, rest)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
