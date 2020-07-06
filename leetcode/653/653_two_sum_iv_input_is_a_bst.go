/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-06 14:45:57
 */

package _653

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	hash := make(map[int]bool)
	return find(root, k, hash)
}

func find(root *TreeNode, k int, hash map[int]bool) bool {
	if root == nil {
		return false
	}

	if hash[k-root.Val] {
		return true
	}

	hash[root.Val] = true
	return find(root.Left, k, hash) || find(root.Right, k, hash)
}
