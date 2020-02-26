/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:25'
*/

package _114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	p := root
	if p.Left == nil {
		return
	}

	p = p.Left
	for p.Right != nil {
		p = p.Right
	}

	p.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
