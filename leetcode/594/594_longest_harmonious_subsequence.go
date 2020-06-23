/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-23 13:47:48
 */

package _594

import "math"

// 枚举
func findLHS(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		count := 0
		flag := false
		for j := 0; j < len(nums); j++ {
			if nums[j] == nums[i] {
				count++
			} else if nums[j]+1 == nums[i] {
				count++
				flag = true
			}
		}

		if flag {
			res = int(math.Max(float64(count), float64(res)))
		}
	}

	return res
}

// 哈希
func findLHS1(nums []int) int {
	hash := make(map[int]int)
	res := 0
	for _, num := range nums {
		hash[num]++
	}

	for k, v := range hash {
		if val, ok := hash[k+1]; ok {
			res = int(math.Max(float64(res), float64(v+val)))
		}
	}

	return res
}
