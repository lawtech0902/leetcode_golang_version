/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-22 13:49:09
 */

package _589

type Node struct {
	Val      int
	Children []*Node
}

// 迭代
func preorder(root *Node) []int {
	var (
		stack []*Node
		res   []int
	)

	if root == nil {
		return res
	}

	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return res
}

// 递归
func preorder1(root *Node) []int {
	var (
		res []int
		dfs func(root *Node)
	)

	dfs = func(root *Node) {
		if root != nil {
			res = append(res, root.Val)
			for _, node := range root.Children {
				dfs(node)
			}
		}
	}

	dfs(root)
	return res
}
