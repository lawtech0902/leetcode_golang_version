package _494

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 上午12:21'
*/

/*
该问题可以转换为 Subset Sum 问题，从而使用 0-1 背包的方法来求解。
可以将这组数看成两部分，P 和 N，其中 P 使用正号，N 使用负号，有以下推导：
                  sum(P) - sum(N) = target
sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
                       2 * sum(P) = target + sum(nums)
因此只要找到一个子集，令它们都取正号，并且和等于 (target + sum(nums))/2，就证明存在解。
 */

func findTargetSumWays(nums []int, S int) int {
	sums := 0
	for _, num := range nums {
		sums += num
	}

	if sums < S || (sums+S)%2 == 1 {
		return 0
	}

	return partitionSubsetSum(nums, (S+sums)/2)
}

func partitionSubsetSum(nums []int, target int) int {
	// dp[i] 表示 nums 中和为 i 的子集个数
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
