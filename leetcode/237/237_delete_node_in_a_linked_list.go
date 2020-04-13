/*
__author__ = 'lawtech'
__date__ = '2020/04/13 10:53 上午'
*/

package _237

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
