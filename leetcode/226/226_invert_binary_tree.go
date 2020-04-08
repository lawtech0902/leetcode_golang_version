/*
__author__ = 'lawtech'
__date__ = '2020/04/08 10:10 上午'
*/

package _226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 迭代
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		cur.Left, cur.Right = cur.Right, cur.Left
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return root
}
