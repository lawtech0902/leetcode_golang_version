package _110

import (
	"go_projects/leetcode_golang_version"
	"math"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午11:03'
*/

type TreeNode = leetcode_golang_version.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)

	if math.Abs(float64(left-right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(getDepth(root.Left)), float64(getDepth(root.Right))))
}
