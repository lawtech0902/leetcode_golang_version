/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午11:17'
*/

package _530

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	nums := inorder(root, []int{})
	min := math.MaxInt32

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < min {
			min = diff
		}
	}

	return min
}

func inorder(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	nums = inorder(root.Left, nums)
	nums = append(nums, root.Val)
	nums = inorder(root.Right, nums)

	return nums
}
