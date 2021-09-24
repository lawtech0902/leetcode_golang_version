package lead_to_offer_ii

import "math"

func minSubArrayLen(target int, nums []int) int {
	min := math.MaxInt32
	sum := 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for i <= j && sum >= target {
			min = int(math.Min(float64(min), float64(j-i+1)))
			sum -= nums[i]
			i++
		}
	}

	if min == math.MaxInt32 {
		return 0
	}

	return min
}
