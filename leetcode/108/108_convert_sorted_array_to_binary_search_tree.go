/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:23'
*/

package _108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)

	if size == 0 {
		return nil
	}

	if size == 1 {
		return &TreeNode{Val: nums[0]}
	}

	root := &TreeNode{Val: nums[size/2]}
	root.Left = sortedArrayToBST(nums[:size/2])
	root.Right = sortedArrayToBST(nums[size/2+1:])

	return root
}
