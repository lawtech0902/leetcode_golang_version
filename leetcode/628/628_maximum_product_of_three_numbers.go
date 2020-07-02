/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-01 16:49:38
 */

package _628

import "math"

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	for _, num := range nums {
		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num <= min2 {
			min2 = num
		}

		if num >= max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max3 = max2
			max2 = num
		} else if num >= max3 {
			max3 = num
		}
	}

	return int(math.Max(float64(min1*min2*max1), float64(max1*max2*max3)))
}
