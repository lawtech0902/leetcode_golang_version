package _108

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:23'
*/

type ListNode = leetcode_golang_version.ListNode
type TreeNode = leetcode_golang_version.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return &TreeNode{Val: nums[0]}
	}

	root := &TreeNode{Val: nums[size/2]}
	root.Left = sortedArrayToBST(nums[:size/2])
	root.Right = sortedArrayToBST(nums[size/2+1:])
	return root
}