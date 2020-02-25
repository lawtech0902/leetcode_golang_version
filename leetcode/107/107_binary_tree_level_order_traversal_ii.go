/*
__author__ = 'lawtech'
__date__ = '2018/8/19 上午1:29'
*/

package _107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(*TreeNode, int)
	)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的level
		if level >= len(res) {
			res = append([][]int{[]int{}}, res...)
		}

		n := len(res)
		res[n-level-1] = append(res[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
