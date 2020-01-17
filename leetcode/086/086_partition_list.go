/*
__author__ = 'lawtech'
__date__ = '2020/1/17 3:41 下午'
*/

package _86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	head1 := &ListNode{Val: 0}
	head2 := &ListNode{Val: 0}
	p1 := head1
	p2 := head2

	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}

		p = p.Next
	}

	p1.Next = head2.Next
	p2.Next = nil
	return head1.Next
}
