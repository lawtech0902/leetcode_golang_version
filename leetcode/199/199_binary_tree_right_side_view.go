/*
__author__ = 'lawtech'
__date__ = '2020/03/23 10:56 上午'
*/

package _199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	// bfs 层序遍历，保留每层最后一个节点
	var (
		res   []int
		queue []*TreeNode
	)

	if root == nil {
		return res
	}

	queue = append(queue, root)
	for len(queue) != 0 {
		curLen := len(queue)
		for curLen != 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			curLen--
			if curLen == 0 {
				res = append(res, cur.Val)
			}
		}
	}

	return res
}
