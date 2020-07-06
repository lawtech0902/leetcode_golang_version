/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-06 15:36:58
 */

package _655

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	d := depth(root)
	res := make([][]string, d)
	cols := int(math.Pow(2, float64(d)) - 1)
	for i := range res {
		res[i] = make([]string, cols)
	}

	populate(root, 0, 0, len(res[0]), res)
	return res
}

func populate(root *TreeNode, row, low, high int, res [][]string) {
	if root == nil {
		return
	}

	mid := (low + high) / 2
	res[row][mid] = strconv.Itoa(root.Val)
	populate(root.Left, row+1, low, mid, res)
	populate(root.Right, row+1, mid, high, res)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left)
	r := depth(root.Right)
	if l > r {
		return l + 1
	}

	return r + 1
}
