/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-23 14:55:14
 */

package _599

import "math"

func findRestaurant(list1, list2 []string) []string {
	var res []string
	hash := make(map[string]int, 0)

	for i, v := range list1 {
		hash[v] = i
	}

	minSum, sum := math.MaxInt32, 0
	for j := 0; j < len(list2) && j <= minSum; j++ {
		if val, ok := hash[list2[j]]; ok {
			sum = j + val
			if sum < minSum {
				res = nil
				res = append(res, list2[j])
				minSum = sum
			} else if minSum == sum {
				res = append(res, list2[j])
			}
		}
	}

	return res
}
