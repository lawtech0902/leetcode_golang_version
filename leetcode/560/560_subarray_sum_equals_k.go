/*
__author__ = 'lawtech'
__date__ = '2020/06/17 10:02 上午'
*/

package _560

// 枚举
func subarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// 前缀表+哈希表优化
func subarraySum1(nums []int, k int) int {
	count, pre := 0, 0
	hash := make(map[int]int)
	hash[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := hash[pre-k]; ok {
			count += hash[pre-k]
		}

		hash[pre] += 1
	}

	return count
}
