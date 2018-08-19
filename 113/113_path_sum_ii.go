package _113

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午4:50'
*/

type TreeNode = leetcode_golang_version.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	var (
		res  [][]int
		path []int
		dfs  func(*TreeNode, int, int)
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
