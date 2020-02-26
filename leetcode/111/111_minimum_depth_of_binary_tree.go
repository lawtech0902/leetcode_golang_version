/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午10:57'
*/

package _111

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	}

	return 1 + int(math.Min(float64(left), float64(right)))
}
