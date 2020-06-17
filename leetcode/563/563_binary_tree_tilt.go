/*
__author__ = 'lawtech'
__date__ = '2020/06/17 10:48 上午'
*/

package _563

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var (
		res        int
		subTreeSum func(root *TreeNode) int
	)

	subTreeSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lSum := subTreeSum(root.Left)
		rSum := subTreeSum(root.Right)
		res += int(math.Abs(float64(lSum - rSum)))
		return root.Val + lSum + rSum
	}

	subTreeSum(root)
	return res
}
