package lead_to_offer_ii

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 2 {
		return 0
	}

	i, j, res, product := 0, 0, 0, 1
	for j < len(nums) {
		product *= nums[j]
		for product >= k {
			product /= nums[i]
			i++
		}

		res += j - i + 1
		j++
	}

	return res
}
