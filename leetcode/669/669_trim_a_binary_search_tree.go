/*
__author__ = 'lawtech'
__date__ = '2020/07/08 2:51 下午'
*/

package _669

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L, R int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > R {
		return trimBST(root.Left, L, R)
	}

	if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
