/*
__author__ = 'lawtech'
__date__ = '2020/03/03 6:45 下午'
*/

package _129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var helper func(root *TreeNode, preSum int) int

	helper = func(root *TreeNode, preSum int) int {
		if root == nil {
			return 0
		}

		preSum = preSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return preSum
		}

		return helper(root.Left, preSum) + helper(root.Right, preSum)
	}

	return helper(root, 0)
}
