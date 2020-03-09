/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午11:09'
*/

package _145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func postorderTraversal(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 迭代
//func postorderTraversal(root *TreeNode) []int {
//	stack := []*TreeNode{root}
//	var res []int
//
//	for len(stack) > 0 {
//		node := stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//
//		if node != nil {
//			res = append(res, node.Val)
//
//			if node.Left != nil {
//				stack = append(stack, node.Left)
//			}
//
//			if node.Right != nil {
//				stack = append(stack, node.Right)
//			}
//		}
//	}
//
//	// reverse
//	i, j := 0, len(res)-1
//	for i < j {
//		res[i], res[j] = res[j], res[i]
//		i++
//		j--
//	}
//
//	return res
//}
