/*
__author__ = 'lawtech'
__date__ = '2020/03/19 9:53 上午'
*/

package _189

// 三种方法

func rotate1(nums []int, k int) {
	// 暴力，O(n * k), O(1)
	var temp, prev int

	size := len(nums)
	for i := 0; i < k; i++ {
		prev = nums[size-1]
		for j := 0; j < size; j++ {
			temp = nums[j]
			nums[j] = prev
			prev = temp
		}
	}
}

func rotate2(nums []int, k int) {
	// 环状替换, O(n), O(1)
	start := 0
	for count := 0; count < len(nums); {
		current := start
		tmp := 0

		for {
			count++
			nums[current], tmp = tmp, nums[current]
			current = (current + k) % len(nums)

			if current == start {
				nums[current], tmp = tmp, nums[current]
				break
			}
		}

		start++
	}
}

func rotate(nums []int, k int) {
	// 反转, O(n), O(1)
	size := len(nums)
	k %= len(nums)
	reverse(nums, 0, size-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, size-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
