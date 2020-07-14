/*
__author__ = 'lawtech'
__date__ = '2020/07/14 2:34 下午'
*/

package _701

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

// 迭代
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		if val > node.Val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				return root
			} else {
				node = node.Right
			}
		} else {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				return root
			} else {
				node = node.Left
			}
		}
	}

	return &TreeNode{Val: val}
}
