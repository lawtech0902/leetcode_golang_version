/*
__author__ = 'lawtech'
__date__ = '2020/2/4 2:36 下午'
*/

package _95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode
	if n == 0 {
		return res
	}

	return dfs(1, n)
}

func dfs(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var res []*TreeNode
	for rootVal := start; rootVal < end+1; rootVal++ {
		leftTree := dfs(start, rootVal-1)
		rightTree := dfs(rootVal+1, end)

		for _, left := range leftTree {
			for _, right := range rightTree {
				root := &TreeNode{
					Val:   rootVal,
					Left:  left,
					Right: right,
				}

				res = append(res, root)
			}
		}
	}

	return res
}
