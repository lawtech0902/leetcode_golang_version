/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午11:02'
*/

package _94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var (
		stack []*TreeNode
		res   []int
	)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		size := len(stack)

		if size == 0 {
			return res
		}

		node := stack[size-1]
		stack = stack[:size-1]

		res = append(res, node.Val)
		root = node.Right
	}
}
