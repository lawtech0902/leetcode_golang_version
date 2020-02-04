/*
__author__ = 'lawtech'
__date__ = '2020/2/4 2:54 下午'
*/

package _99

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var (
		firstNode, secondNode, pre, lastNode *TreeNode
	)

	node := root
	for node != nil {
		// 存在子左树
		if node.Left != nil {

			// 查找当前节点的前驱
			pre = node.Left
			for pre.Right != nil && pre.Right != node {
				pre = pre.Right
			}

			// 找到前驱节点, 记录后驱为当前节点
			if pre.Right == nil {
				pre.Right = node
				node = node.Left
				continue
			}

			// 恢复树结构
			pre.Right = nil
		}

		// 存放错误节点
		if lastNode != nil && lastNode.Val > node.Val {
			if firstNode == nil {
				firstNode = lastNode
				secondNode = node
			} else {
				secondNode = node
			}
		}

		lastNode, node = node, node.Right
	}

	if secondNode != nil {
		firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
	}
}
