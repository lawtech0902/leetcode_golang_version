/*
__author__ = 'lawtech'
__date__ = '2020/07/30 2:00 下午'
*/

package _28

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
