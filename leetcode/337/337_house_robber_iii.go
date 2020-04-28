/*
__author__ = 'lawtech'
__date__ = '2020/04/28 9:48 上午'
*/

package _337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(robber(root))
}

func robber(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	money := root.Val
	l1, l2 := robber(root.Left)
	r1, r2 := robber(root.Right)
	c1 := money + l2 + r2
	c2 := max(l1, l2) + max(r1, r2)
	return c1, c2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
