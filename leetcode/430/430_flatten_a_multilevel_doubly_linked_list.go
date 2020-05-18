/*
__author__ = 'lawtech'
__date__ = '2020/05/18 1:44 下午'
*/

package _430

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	dummy := &Node{
		Val:   0,
		Prev:  nil,
		Next:  root,
		Child: nil,
	}

	dfs(dummy, root)
	dummy.Next.Prev = nil
	return dummy.Next
}

func dfs(prev, curr *Node) *Node {
	if curr == nil {
		return prev
	}

	curr.Prev = prev
	prev.Next = curr

	tempNext := curr.Next
	tail := dfs(curr, curr.Child)
	curr.Child = nil
	return dfs(tail, tempNext)
}
