/*
__author__ = 'lawtech'
__date__ = '2020/08/04 5:31 下午'
*/

package _46

import "strconv"

// base dp O(n) O(n)
func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	size := len(numStr)
	dp := make([]int, size+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= len(numStr); i++ {
		dp[i] = dp[i-1]
		if temp, _ := strconv.Atoi(numStr[i-2 : i]); temp >= 10 && temp <= 25 {
			dp[i] += dp[i-2]
		}
	}

	return dp[size]
}

// O(n) O(1)
func translateNum1(num int) int {
	a, b := 1, 1
	x, y := 0, num%10
	for num != 0 {
		num /= 10
		x = num % 10
		tmp := 10*x + y
		c := 0
		if tmp >= 10 && tmp <= 25 {
			c = a + b
		} else {
			c = a
		}

		b = a
		a = c
		y = x
	}

	return a
}
