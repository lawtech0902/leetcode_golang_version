package leetcode_golang_version

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:50'
*/

func LengthOfLCS(nums1, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)

	// 傻逼一样的二维数组声明
	dp := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 傻逼一样的求最大值
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return dp[n1][n2]
}
