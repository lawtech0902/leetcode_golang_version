package lead_to_offer_ii

import "math"

func divide(a, b int) int {
	sign := 1
	if a^b < 0 {
		sign = -1
	}

	if a < 0 {
		a = ^a + 1
	}

	if b < 0 {
		b = ^b + 1
	}

	res := 0
	for i := 31; i >= 0; i-- {
		if a>>i >= b {
			if i == 31 && sign == 1 {
				return math.MaxInt32
			}

			res += 1 << i
			a -= b << i
		}
	}

	if sign == -1 {
		return ^res + 1
	}

	return res
}
