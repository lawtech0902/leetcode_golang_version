package lead_to_offer_ii

func singleNumbers(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range nums {
			count += num >> i & 1
		}

		if count%3 == 1 {
			if i == 31 {
				res -= 1 << i
			} else {
				res |= 1 << i
			}
		}
	}

	return res
}

func singleNumbers1(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = ^two & (one ^ num)
		two = ^one & (two ^ num)
	}

	return one
}
