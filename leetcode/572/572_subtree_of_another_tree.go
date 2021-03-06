/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午3:20'
*/

package _572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val && isSame(s, t) {
		return true
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
