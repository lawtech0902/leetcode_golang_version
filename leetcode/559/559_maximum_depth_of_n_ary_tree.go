/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-17 09:36:23
 */

package _559

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	m := 1
	for _, v := range root.Children {
		m = max(m, maxDepth(v)+1)
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
