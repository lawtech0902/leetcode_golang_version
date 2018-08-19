package leetcode_golang_version

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:28'
*/

// leetcode没有
func convert(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	convert(root.Left)
	left := root.Left

	// 连接root与左子树最大节点
	if left != nil {
		for left.Right != nil {
			left = left.Right
		}

		root.Left, left.Right = left.Right, root.Left
	}

	convert(root.Right)
	right := root.Right

	// 连接root与右子树最小节点
	if right != nil {
		for right.Left != nil {
			right = right.Left
		}

		root.Right, right.Left = right.Left, root.Right
	}

	for root.Left != nil {
		root = root.Left
	}

	return root
}
