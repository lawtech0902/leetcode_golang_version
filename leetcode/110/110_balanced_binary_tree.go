/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午11:03'
*/

package _110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
