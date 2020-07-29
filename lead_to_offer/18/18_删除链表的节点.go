/*
__author__ = 'lawtech'
__date__ = '2020/07/29 1:57 下午'
*/

package _18

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	pre, cur := head, head.Next
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}

	if cur != nil {
		pre.Next = cur.Next
	}

	return head
}
