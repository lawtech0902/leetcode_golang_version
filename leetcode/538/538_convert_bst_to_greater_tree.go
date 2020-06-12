/*
__author__ = 'lawtech'
__date__ = '2020/06/12 5:28 下午'
*/

package _538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var (
		total    int
		traverse func(root *TreeNode)
	)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Right)
		root.Val += total
		total = root.Val
		traverse(root.Left)
	}

	traverse(root)
	return root
}
