/*
__author__ = 'lawtech'
__date__ = '2020/05/28 10:25 上午'
*/

package _477

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
