/*
__author__ = 'lawtech'
__date__ = '2020/05/28 10:25 上午'
*/

package _477

// 考虑数组中每个数二进制的第 i 位，假设一共有 t 个 0 和 n - t 个 1，那么显然在第 i 位的汉明距离的总和为 t * (n - t)。
func totalHammingDistance(nums []int) int {
	res, n := 0, len(nums)
	for i := 0; i < 32; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			cnt += (nums[j] >> i) & 1
		}

		res += (n - cnt) * cnt
	}

	return res
}
