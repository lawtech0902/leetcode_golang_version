/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-07 11:03:23
 */

package _662

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	cur := []*TreeNode{
		{0, root.Left, root.Right},
	}
	next := []*TreeNode{}
	max := 0
	for len(cur) != 0 {
		length := cur[len(cur)-1].Val - cur[0].Val + 1
		if length > max {
			max = length
		}

		for _, node := range cur {
			if node.Left != nil {
				next = append(next, &TreeNode{node.Val * 2, node.Left.Left, node.Left.Right})
			}

			if node.Right != nil {
				next = append(next, &TreeNode{node.Val*2 + 1, node.Right.Left, node.Right.Right})
			}
		}

		cur = next
		next = []*TreeNode{}
	}

	return max
}
