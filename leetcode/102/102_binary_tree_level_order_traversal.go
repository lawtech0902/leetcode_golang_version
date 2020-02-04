/*
__author__ = 'lawtech'
__date__ = '2018/8/19 上午1:01'
*/

package _102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(*TreeNode, int)
	)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
