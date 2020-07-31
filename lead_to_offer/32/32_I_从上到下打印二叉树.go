/*
__author__ = 'lawtech'
__date__ = '2020/07/31 10:19 上午'
*/

package _32

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := list.New()
	queue.PushBack(root)
	var res []int
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		res = append(res, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return res
}

// 用这个空间效率更高
func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	var res []int
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
}
