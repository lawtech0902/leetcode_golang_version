/*
__author__ = 'lawtech'
__date__ = '2020/07/14 2:04 下午'
*/

package _700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

// 迭代
func searchBST1(root *TreeNode, val int) *TreeNode {
	for root != nil && val != root.Val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root
}
