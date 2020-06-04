/*
__author__ = 'lawtech'
__date__ = '2020/06/04 2:21 下午'
*/

package _515

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		size := len(stack)
		levelMax := math.MinInt64
		for i := 0; i < size; i++ {
			curNode := stack[0]
			stack = stack[1:]
			levelMax = max(curNode.Val, levelMax)
			if curNode.Left != nil {
				stack = append(stack, curNode.Left)
			}

			if curNode.Right != nil {
				stack = append(stack, curNode.Right)
			}
		}

		res = append(res, levelMax)
	}

	return res
}

// dfs
func largestValues1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append(res, math.MinInt64)
		}

		res[level] = max(res[level], root.Val)
		if root.Left != nil {
			dfs(root.Left, level+1)
		}

		if root.Right != nil {
			dfs(root.Right, level+1)
		}
	}

	dfs(root, 0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
