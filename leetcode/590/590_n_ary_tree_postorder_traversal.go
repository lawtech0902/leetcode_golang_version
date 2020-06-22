/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-22 14:47:47
 */

package _590

type Node struct {
	Val      int
	Children []*Node
}

// 迭代
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := make([]*Node, 0)
	stack = append(stack, root)
	res := make([]int, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{node.Val}, res...)
		for i := 0; i < len(node.Children); i++ {
			stack = append(stack, node.Children[i])
		}
	}

	return res
}

// 递归
func postorder1(root *Node) []int {
	var res []int

	if root == nil {
		return res
	}

	for _, node := range root.Children {
		res = append(res, postorder1(node)...)
	}

	res = append(res, root.Val)
	return res
}
