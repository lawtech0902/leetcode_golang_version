package _144

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午10:46'
*/

type TreeNode = leetcode_golang_version.TreeNode

func preorderTraversal(root *TreeNode) []int {
	// 迭代
	stack := []*TreeNode{root}
	var res []int

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			res = append(res, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return res
}
