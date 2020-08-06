/*
__author__ = 'lawtech'
__date__ = '2020/08/06 2:09 下午'
*/

package _55

import "math"

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

	return 1 + max(getDepth(root.Left), getDepth(root.Right))
}
