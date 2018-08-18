package _94

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午11:02'
*/

type TreeNode = leetcode_golang_version.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var (
		stack []*TreeNode
		res   []int
	)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		size := len(stack)

		if size == 0 {
			return res
		}

		node := stack[size-1]
		stack = stack[:size-1]

		res = append(res, node.Val)
		root = node.Right
	}
}
