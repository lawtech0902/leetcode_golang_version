/*
__author__ = 'lawtech'
__date__ = '2020/07/30 1:39 下午'
*/

package _27

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	leftRoot := mirrorTree(root.Right)
	rightRoot := mirrorTree(root.Left)
	root.Left = leftRoot
	root.Right = rightRoot
	return root
}

// 迭代
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[0]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		node.Left, node.Right = node.Right, node.Left
		stack = stack[1:]
	}

	return root
}
