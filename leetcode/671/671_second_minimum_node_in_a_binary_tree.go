/*
__author__ = 'lawtech'
__date__ = '2020/07/08 3:59 下午'
*/

package _671

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}

	left, right := root.Left.Val, root.Right.Val
	if left == root.Val {
		left = findSecondMinimumValue(root.Left)
	}

	if right == root.Val {
		right = findSecondMinimumValue(root.Right)
	}

	if left == -1 || right == -1 {
		return max(left, right)
	}

	return min(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
