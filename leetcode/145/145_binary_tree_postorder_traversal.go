package _145

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午11:09'
*/

type TreeNode = leetcode_golang_version.TreeNode

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	var res []int

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			res = append(res, node.Val)

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
