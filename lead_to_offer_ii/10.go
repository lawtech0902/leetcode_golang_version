package lead_to_offer_ii

func subarrySum(nums []int, k int) int {
	var res, sum int
	hash := make(map[int]int)
	hash[0]++
	for _, num := range nums {
		sum += num
		res += hash[sum-k]
		hash[sum]++
	}

	return res
}
