package _111

import (
	"go_projects/leetcode_golang_version"
	"math"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午10:57'
*/

type TreeNode = leetcode_golang_version.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}

	return 1 + int(math.Min(float64(left), float64(right)))
}
