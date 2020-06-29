/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-29 15:17:15
 */

package _617

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归，先序遍历
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// 迭代
func mergeTrees1(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	stack := list.New()
	stack.PushBack([]*TreeNode{t1, t2})
	for stack.Len() != 0 {
		t := stack.Back().Value.([]*TreeNode)
		stack.Remove(stack.Back())
		if t[0] == nil || t[1] == nil {
			continue
		}

		t[0].Val += t[1].Val
		if t[0].Left == nil {
			t[0].Left = t[1].Left
		} else {
			stack.PushBack([]*TreeNode{t[0].Left, t[1].Left})
		}

		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			stack.PushBack([]*TreeNode{t[0].Right, t[1].Right})
		}
	}

	return t1
}
