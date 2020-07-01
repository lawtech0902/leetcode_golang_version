/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-01 11:08:21
 */

package _623

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return root
	}

	if d == 1 {
		node := &TreeNode{Val: v}
		node.Left = root
		return node
	}

	if d == 2 {
		l := &TreeNode{Val: v}
		r := &TreeNode{Val: v}

		l.Left = root.Left
		r.Right = root.Right

		root.Left = l
		root.Right = r
		return root
	}

	addOneRow(root.Left, v, d-1)
	addOneRow(root.Right, v, d-1)
	return root
}
