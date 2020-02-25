/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午11:19'
*/

package _105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
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
