/*
__author__ = 'lawtech'
__date__ = '2020/03/02 9:14 下午'
*/

package _124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var (
		res         int
		_maxPathSum func(root *TreeNode) int
	)

	res = math.MinInt32

	_maxPathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		sum := root.Val
		lMax, rMax := 0, 0
		if root.Left != nil {
			lMax = _maxPathSum(root.Left)
			if lMax > 0 {
				sum += lMax
			}
		}

		if root.Right != nil {
			rMax = _maxPathSum(root.Right)
			if rMax > 0 {
				sum += rMax
			}
		}

		res = int(math.Max(float64(res), float64(sum)))
		return int(math.Max(float64(root.Val), math.Max(float64(root.Val+lMax), float64(root.Val+rMax))))
	}

	_maxPathSum(root)
	return res
}
