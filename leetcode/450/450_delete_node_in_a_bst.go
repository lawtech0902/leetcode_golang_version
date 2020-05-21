/*
__author__ = 'lawtech'
__date__ = '2020/05/21 1:48 下午'
*/

package _450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}

	return root
}

func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}

	return root.Val
}

func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Left != nil {
		root = root.Right
	}

	return root.Val
}
