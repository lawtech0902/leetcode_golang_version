package _114

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:25'
*/

type TreeNode = leetcode_golang_version.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	p := root
	if p.Left == nil {
		return
	}
	p = p.Left
	for p.Right != nil {
		p = p.Right
	}

	p.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
