/*
__author__ = 'lawtech'
__date__ = '2020/08/03 10:14 上午'
*/

package _34

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var (
		res  [][]int
		path []int
		dfs  func(root *TreeNode, level, sum int)
	)

	dfs = func(root *TreeNode, level, sum int) {
		if root == nil {
			return
		}

		if level >= len(path) {
			path = append(path, root.Val)
		} else {
			path[level] = root.Val
		}

		sum -= root.Val
		if root.Left == nil && root.Right == nil && sum == 0 {
			tmp := make([]int, level+1)
			copy(tmp, path)
			res = append(res, tmp)
		}

		dfs(root.Left, level+1, sum)
		dfs(root.Right, level+1, sum)
	}

	dfs(root, 0, sum)
	return res
}
