/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:38'
*/

package _300

import "math"

// dp O(n2) O(n)
func lengthOfLIS(nums []int) int {
	if nums == nil {
		return 0
	}

	var dp []int
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}

	return max(dp)
}

func max(nums []int) int {
	res := -math.MaxInt32
	for _, num := range nums {
		if res < num {
			res = num
		}
	}

	return res
}

/*
贪心 + 二分
设当前已求出的最长上升子序列的长度为 len（初始时为 11），从前往后遍历数组 nums，在遍历到 nums[i] 时：
如果 nums[i]>d[len] ，则直接加入到 dd 数组末尾，并更新 len=len+1；
否则，在 dd 数组中二分查找，找到第一个比 nums[i] 小的数 d[k]d[k] ，并更新 d[k+1]=nums[i]。
*/
func lengthOfLIS1(nums []int) int {
	var d []int

	for _, num := range nums {
		if d == nil || num > d[len(d)-1] {
			d = append(d, num)
		} else {
			l, r := 0, len(d)-1
			loc := r
			for l <= r {
				mid := l + (r-l)/2
				if d[mid] >= num {
					loc = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}

			d[loc] = num
		}
	}

	return len(d)
}
