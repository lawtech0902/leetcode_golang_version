/*
__author__ = 'lawtech'
__date__ = '2020/06/08 9:50 上午'
*/

package _523

// 记忆化暴力
func checkSubarraySum(nums []int, k int) bool {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	for start := 0; start < len(nums)-1; start++ {
		for end := start + 1; end < len(nums); end++ {
			sum := sums[end] - sums[start] + nums[start]
			if sum == k || (k != 0 && sum%k == 0) {
				return true
			}
		}
	}

	return false
}

func checkSubarraySum1(nums []int, k int) bool {
	sum := 0
	hash := make(map[int]int)
	hash[0] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}

		if val, ok := hash[sum]; ok {
			if i-val > 1 {
				return true
			}
		} else {
			hash[sum] = i
		}
	}

	return false
}
