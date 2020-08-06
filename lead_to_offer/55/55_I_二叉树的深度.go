/*
__author__ = 'lawtech'
__date__ = '2020/08/06 2:02 下午'
*/

package _55

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(values ...int) int {
	maxValue := math.MinInt32
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}
