package lead_to_offer_ii

// 滑动窗口 + hash
func findAnagrams(s, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}

	hash := make([]int, 26)
	for i := range p {
		hash[p[i]-'a']++
	}

	left, right := 0, 0
	count := len(p)
	for right < len(s) {
		if hash[s[right]-'a'] > 0 {
			count--
		}

		hash[s[right]-'a']--
		right++
		if count == 0 {
			res = append(res, left)
		}

		if right-left == len(p) {
			if hash[s[left]-'a'] >= 0 {
				count++
			}

			hash[s[left]-'a']++
			left++
		}
	}

	return res
}
