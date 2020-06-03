/*
__author__ = 'lawtech'
__date__ = '2020/06/03 10:52 上午'
*/

package _501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bst中序有序
func findMode(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	m := make(map[int]int)
	var (
		max int
		res []int
	)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		m[node.Val]++
		if max < m[node.Val] {
			max = m[node.Val]
			res = nil
			res = append(res, node.Val)
		} else if max == m[node.Val] {
			res = append(res, node.Val)
		}

		root = node.Right
	}

	return res
}
