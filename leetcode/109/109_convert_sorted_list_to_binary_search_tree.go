package _109

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:12'
*/

type ListNode = leetcode_golang_version.ListNode
type TreeNode = leetcode_golang_version.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	p := head

	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}

	return sortedArrToBST(nums)
}

func sortedArrToBST(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return &TreeNode{Val: nums[0]}
	}

	root := &TreeNode{Val: nums[size/2]}
	root.Left = sortedArrToBST(nums[:size/2])
	root.Right = sortedArrToBST(nums[size/2+1:])
	return root
}
