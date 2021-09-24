package lead_to_offer_ii

func findMaxLength(nums []int) int {
	var res, cnt int
	hash := make(map[int]int)
	hash[0] = -1
	for i, num := range nums {
		if num == 0 {
			cnt--
		} else {
			cnt++
		}

		if v, ok := hash[cnt]; ok {
			res = max(res, i-v)
		} else {
			hash[cnt] = i
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
