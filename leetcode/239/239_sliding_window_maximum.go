/*
__author__ = 'lawtech'
__date__ = '2020/04/13 1:27 下午'
*/

package _239

// 暴力 O(Nk)
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	n := len(nums)
	if n*k == 0 {
		return res
	}

	for i := 0; i < n-k+1; i++ {
		res = append(res, max(nums[i:i+k]))
	}

	return res
}

func max(nums []int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

// 双向队列 O(N)
func maxSlidingWindow1(nums []int, k int) []int {
	var (
		res, window []int
	)

	for i := 0; i < len(nums); i++ {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}

		window = append(window, i)

		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}

	return res
}
