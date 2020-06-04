/*
__author__ = 'lawtech'
__date__ = '2020/06/04 10:57 上午'
*/

package _508

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	var (
		res []int
		max int
		dfs func(root *TreeNode) int
	)

	hash := make(map[int]int)

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		cur := left + right + root.Val
		hash[cur]++
		if hash[cur] > max {
			res = []int{cur}
			max = hash[cur]
		} else if hash[cur] == max {
			res = append(res, cur)
		}

		return cur
	}

	dfs(root)
	return res
}
