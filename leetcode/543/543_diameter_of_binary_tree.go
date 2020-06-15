/*
__author__ = 'lawtech'
__date__ = '2020/06/15 3:00 下午'
*/

package _543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var (
		res int
		dfs func(node *TreeNode) int
	)

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := dfs(node.Left)
		r := dfs(node.Right)
		res = max(res, l+r)
		return 1 + max(l, r)
	}

	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
