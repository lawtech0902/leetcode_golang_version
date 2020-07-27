/*
__author__ = 'lawtech'
__date__ = '2020/07/27 1:56 下午'
*/

package _7

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preoder + inorder => tree
// 递归
func buildTree(preorder, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	if len(inorder) == 1 {
		return root
	}

	idx := getIndex(root.Val, inorder)
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

func getIndex(val int, nums []int) int {
	for i, num := range nums {
		if num == val {
			return i
		}
	}

	return 0
}

// 迭代
func buildTree1(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	size := len(preorder)
	stack := list.New()
	stack.PushBack(root)
	inorderIndex := 0
	for i := 1; i < size; i++ {
		preorderVal := preorder[i]
		node := stack.Back().Value.(*TreeNode)
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{Val: preorderVal}
			stack.PushBack(node.Left)
		} else {
			for stack.Len() != 0 && stack.Back().Value.(*TreeNode).Val == inorder[inorderIndex] {
				node = stack.Back().Value.(*TreeNode)
				stack.Remove(stack.Back())
				inorderIndex++
			}

			node.Right = &TreeNode{Val: preorderVal}
			stack.PushBack(node.Right)
		}
	}

	return root
}
