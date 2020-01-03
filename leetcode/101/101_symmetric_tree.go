package _101

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午3:27'
*/

type TreeNode = leetcode_golang_version.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSame(root.Left, root.Right)
}

func isSame(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return isSame(l.Left, r.Right) && isSame(l.Right, r.Left)
}
