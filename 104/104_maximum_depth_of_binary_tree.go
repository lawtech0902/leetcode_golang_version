package _104

import (
	"go_projects/leetcode_golang_version"
	"math"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午10:52'
*/

type TreeNode = leetcode_golang_version.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}
