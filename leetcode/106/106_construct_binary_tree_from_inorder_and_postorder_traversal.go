package _106

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午11:36'
*/

type TreeNode = leetcode_golang_version.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	if len(inorder) == 1 {
		return root
	}

	idx := getIndex(root.Val, inorder)

	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	return root
}

func getIndex(val int, nums []int) int {
	for i, num := range nums {
		if num == val {
			return i
		}
	}
	return 0
}
