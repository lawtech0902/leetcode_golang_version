/*
__author__ = 'lawtech'
__date__ = '2020/04/01 1:27 下午'
*/

package _222

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归O(n) O(logN)
func countNodes(root *TreeNode) int {
	if root != nil {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	} else {
		return 0
	}
}

// 二分 O(logN * logN) O(1)
// 如果树为空，返回 0。
// 计算树的高度 d。
// 如果 d == 0，返回 1。
// 除最后一层以外的所有节点数为 2^d-1。最后一层的节点数通过二分搜索，检查最后一层有多少个节点。使用函数 exists(idx, d, root) 检查第 idx 节点是否存在。
// 使用二分搜索实现 exists(idx, d, root)。
// 返回 2^d - 1 + 最后一层的节点数。
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d := depth(root)
	if d == 0 {
		return 1
	}

	left, right := 1, int(math.Pow(2, float64(d)))-1
	for left <= right {
		pivot := left + (right-left)/2
		if exists(pivot, d, root) {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return int(math.Pow(2, float64(d))) - 1 + left
}

func exists(idx, d int, root *TreeNode) bool {
	left, right := 1, int(math.Pow(2, float64(d)))-1
	for i := 0; i < d; i++ {
		pivot := left + (right-left)/2
		if idx < pivot {
			root = root.Left
			right = pivot
		} else {
			root = root.Right
			left = pivot + 1
		}
	}

	return root != nil
}

func depth(root *TreeNode) int {
	res := 0
	for root.Left != nil {
		root = root.Left
		res += 1
	}

	return res
}
