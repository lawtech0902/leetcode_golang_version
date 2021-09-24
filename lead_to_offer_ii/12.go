package lead_to_offer_ii

func pivotIndex(nums []int) int {
	sum, cur := 0, 0
	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		sum -= num
		if cur == sum {
			return i
		}

		cur += num
	}

	return -1
}
