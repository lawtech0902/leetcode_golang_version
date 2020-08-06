/*
__author__ = 'lawtech'
__date__ = '2020/08/06 11:26 上午'
*/

package _54

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var tmp []int
	inorder(root, &tmp)
	return tmp[k-1]
}

// 递归，中序有序，中序倒序求kthlargest
func inorder(root *TreeNode, tmp *[]int) {
	if root == nil {
		return
	}

	inorder(root.Right, tmp)
	*tmp = append(*tmp, root.Val)
	inorder(root.Left, tmp)
}

// 迭代
func kthLargest1(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}

		root = root.Left
	}
}
