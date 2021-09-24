package lead_to_offer_ii

// 滑动窗口 + hash
func lengthOfLongestSubstring(s string) int {
	hash := make([]int, 128)
	for i := range hash {
		hash[i] = -1
	}

	res, left := 0, 0
	for i, byte := range s {
		if hash[byte] >= left {
			left = hash[byte] + 1
		} else if i-left+1 > res {
			res = i - left + 1
		}

		hash[byte] = i
	}

	return res
}
