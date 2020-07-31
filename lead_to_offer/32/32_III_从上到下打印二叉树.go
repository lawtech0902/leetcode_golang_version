/*
__author__ = 'lawtech'
__date__ = '2020/07/31 2:03 下午'
*/

package _32

// dfs
func levelOrderIII(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(root *TreeNode, level int)
	)

	dfs = func(root *TreeNode, level int) {
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
