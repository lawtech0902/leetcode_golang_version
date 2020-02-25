/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午5:12'
*/

package _109

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	p := head

	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}

	return sortedArrToBST()
}

func sortedArrToBST(nums []int) *TreeNode {
	size := len(nums)

	if size == 0 {
		return nil
	}

	if size == 1 {
		return &TreeNode{Val: nums[0]}
	}

	root := &TreeNode{Val: nums[size/2]}
	root.Left = sortedArrToBST(nums[:size/2])
	root.Right = sortedArrToBST(nums[size/2+1:])

	return root
}
