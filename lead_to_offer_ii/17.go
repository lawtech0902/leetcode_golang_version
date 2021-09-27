package lead_to_offer_ii

func minWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	hash := make([]int, 128)
	for i := range t {
		hash[t[i]]++
	}

	l, cnt, max, res := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			cnt--
		}

		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}

		if cnt == 0 && max > r+1-l {
			max = r + 1 - l
			res = s[l : r+1]
		}
	}

	return res
}
