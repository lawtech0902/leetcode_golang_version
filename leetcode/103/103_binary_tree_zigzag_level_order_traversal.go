/*
__author__ = 'lawtech'
__date__ = '2020/2/4 3:31 下午'
*/

package _03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(*TreeNode, int)
	)

	dfs = func(root *TreeNode, level int) {
		// preOrder
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			tmp := make([]int, len(res[level])+1)
			tmp[0] = root.Val
			copy(tmp[1:], res[level])
			res[level] = tmp
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
