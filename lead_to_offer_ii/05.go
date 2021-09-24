package lead_to_offer_ii

import "math"

func maxProduct(words []string) int {
	size := len(words)
	arr := make([]int, size)
	for i, word := range words {
		for _, letter := range word {
			arr[i] |= 1 << int(letter-'a')
		}
	}

	res := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if arr[i]&arr[j] == 0 {
				res = int(math.Max(float64(res), float64(len(words[i])*len(words[j]))))
			}
		}
	}

	return res
}
