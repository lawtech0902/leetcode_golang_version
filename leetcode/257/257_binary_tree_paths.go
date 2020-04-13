/*
__author__ = 'lawtech'
__date__ = '2020/04/13 3:47 下午'
*/

package _257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}

	dfs(root, strconv.Itoa(root.Val), &res)
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}

	if root.Right != nil {
		dfs(root.Right, path+"->"+strconv.Itoa(root.Right.Val), res)
	}

	if root.Left != nil {
		dfs(root.Left, path+"->"+strconv.Itoa(root.Left.Val), res)
	}
}
