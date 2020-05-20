/*
__author__ = 'lawtech'
__date__ = '2020/05/19 4:52 下午'
*/

package _437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == sum {
		res += 1
	}

	res += helper(root.Left, sum-root.Val)
	res += helper(root.Right, sum-root.Val)
	return res
}
