/*
__author__ = 'lawtech'
__date__ = '2020/05/27 9:48 上午'
*/

package _473

import "sort"

//dfs
func makesquare(nums []int) bool {
	var dfs func(index int) bool
	if nums == nil || len(nums) == 0 {
		return false
	}

	perimeter := 0
	for _, num := range nums {
		perimeter += num
	}

	possibleSide := perimeter / 4
	if possibleSide*4 != perimeter {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sums := make([]int, 4)

	dfs = func(index int) bool {
		if index == len(nums) {
			return sums[0] == sums[1] && sums[1] == sums[2] && sums[2] == possibleSide
		}

		for i := 0; i < 4; i++ {
			if sums[i]+nums[index] <= possibleSide {
				sums[i] += nums[index]
				if dfs(index + 1) {
					return true
				}

				sums[i] -= nums[index]
			}
		}

		return false
	}

	return dfs(0)
}
