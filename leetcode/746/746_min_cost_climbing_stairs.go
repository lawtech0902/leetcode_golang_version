/*
__author__ = 'lawtech'
__date__ = '2020/07/21 1:46 下午'
*/

package _746

import "math"

// dp
func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := 0; i < len(cost); i++ {
		minCost := cost[i] + min(f1, f2)
		f1 = f2
		f2 = minCost
	}

	return min(f1, f2)
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
