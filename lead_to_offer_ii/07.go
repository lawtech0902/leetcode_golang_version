package lead_to_offer_ii

import "sort"

func threeSum(nums []int) [][]int {
	size := len(nums)
	var res [][]int
	if size < 3 {
		return res
	}

	sort.Ints(nums)
	for i, num := range nums {
		if num > 0 {
			return res
		}

		if i > 0 && num == nums[i-1] {
			continue
		}

		l, r := i+1, size-1
		for l < r {
			if num+nums[l]+nums[r] == 0 {
				res = append(res, []int{num, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if num+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
