/*
__author__ = 'lawtech'
__date__ = '2020/06/04 2:06 下午'
*/

package _513

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[0]
		stack = stack[1:]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}

	return root.Val
}

// dfs
func findBottomLeftValue1(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	var (
		res, maxLevel int
		dfs           func(root *TreeNode, level int)
	)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		dfs(root.Left, level+1)
		if level > maxLevel {
			maxLevel = level
			res = root.Val
		}

		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
