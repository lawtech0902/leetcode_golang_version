/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-02 09:52:10
 */

package _637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	path := make([][2]float64, 0)
	visit(root, &path, 0)
	for _, v := range path {
		res = append(res, v[0]/v[1])
	}

	return res
}

func visit(root *TreeNode, path *[][2]float64, depth int) {
	if root == nil {
		return
	}

	if len(*path) == depth {
		*path = append(*path, [2]float64{})
	}

	(*path)[depth][0] += float64(root.Val)
	(*path)[depth][1]++

	visit(root.Left, path, depth+1)
	visit(root.Right, path, depth+1)
}

// bfs
func averageOfLevels1(root *TreeNode) []float64 {
	res := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var sum float64
		for i := 0; i < size; i++ {
			sum += float64(queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			queue = queue[1:]
		}

		res = append(res, sum/float64(size))
	}

	return res
}
