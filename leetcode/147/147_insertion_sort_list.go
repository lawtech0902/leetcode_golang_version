/*
__author__ = 'lawtech'
__date__ = '2020/03/10 4:51 下午'
*/

package _147

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	cur := head
	for cur.Next != nil {
		if cur.Next.Val < cur.Val {
			pre := dummy
			for pre.Next.Val < cur.Next.Val {
				pre = pre.Next
			}

			tmp := cur.Next
			cur.Next = tmp.Next
			tmp.Next = pre.Next
			pre.Next = tmp
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
