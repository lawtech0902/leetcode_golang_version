/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-29 10:18:10
 */

package _606

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	s := strconv.Itoa(t.Val)
	l := tree2str(t.Left)
	r := tree2str(t.Right)

	if t.Left == nil && t.Right == nil {
		return s
	}

	if t.Right == nil {
		return s + "(" + l + ")"
	}

	return s + "(" + l + ")" + "(" + r + ")"
}
