/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-06 15:10:22
 */

package _654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	idx := maxValIndex(nums)
	root := &TreeNode{Val: nums[idx]}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}

func maxValIndex(nums []int) int {
	res := 0
	for k, v := range nums {
		if v > nums[res] {
			res = k
		}
	}

	return res
}
