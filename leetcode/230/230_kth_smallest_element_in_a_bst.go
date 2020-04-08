/*
__author__ = 'lawtech'
__date__ = '2020/04/08 11:10 上午'
*/

package _230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var tmp []int
	inOrder(root, &tmp)
	return tmp[k-1]
}

// 递归, 二叉搜索树中序有序
func inOrder(root *TreeNode, tmp *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, tmp)
	*tmp = append(*tmp, root.Val)
	inOrder(root.Right, tmp)
}

// 迭代
func kthSmallest1(root *TreeNode, k int) int {
	var stack []*TreeNode

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}
