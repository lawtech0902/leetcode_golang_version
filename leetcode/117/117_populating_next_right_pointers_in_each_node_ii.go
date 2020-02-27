/*
__author__ = 'lawtech'
__date__ = '2020/02/27 3:55 下午'
*/

package _117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	// 建立虚拟节点 cur, dummy
	dummy := &Node{Val: 0}
	cur := dummy
	curLevel := root

	for curLevel != nil {
		if curLevel.Left != nil {
			cur.Next = curLevel.Left
			cur = cur.Next
		}

		if curLevel.Right != nil {
			cur.Next = curLevel.Right
			cur = cur.Next
		}

		curLevel = curLevel.Next // 当前层节点指针前移, 继续处理对应的下一层的节点.
	}

	connect(dummy.Next) // cur 已经将dummy.next -> root.left
	return root
}
