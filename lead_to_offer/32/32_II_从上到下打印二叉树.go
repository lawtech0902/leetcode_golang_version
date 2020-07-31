/*
__author__ = 'lawtech'
__date__ = '2020/07/31 10:48 上午'
*/

package _32

// bfs
func levelOrderII(root *TreeNode) [][]int {
	var (
		queue []*TreeNode
		res   [][]int
	)

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		var tmp []int
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, tmp)
	}

	return res
}

// 递归 dfs
func levelOrderII1(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(root *TreeNode, level int)
	)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
