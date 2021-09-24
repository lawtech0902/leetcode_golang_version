package lead_to_offer_ii

import "reflect"

func checkInclusion(s1, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}

	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	for i := 0; i < n1; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}

	if reflect.DeepEqual(cnt1, cnt2) {
		return true
	}

	left, right := 0, n1
	for right < n2 {
		cnt2[s2[right]-'a']++
		cnt2[s2[left]-'a']--
		if reflect.DeepEqual(cnt1, cnt2) {
			return true
		}

		right++
		left++
	}

	return false
}
