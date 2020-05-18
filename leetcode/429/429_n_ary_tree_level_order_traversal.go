/*
__author__ = 'lawtech'
__date__ = '2020/05/18 11:05 上午'
*/

package _429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	recursive(root, 0, &res)
	return res
}

func recursive(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}

	for _, v := range root.Children {
		recursive(v, level+1, res)
	}

	for len(*res) <= level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)
}
