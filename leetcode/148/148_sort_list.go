/*
__author__ = 'lawtech'
__date__ = '2020/03/10 5:06 下午'
*/

package _148

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并 O(nlogn) O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 寻找中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head1 := head
	head2 := slow.Next
	slow.Next = nil

	// 归并
	head1 = sortList(head1)
	head2 = sortList(head2)
	head = merge(head1, head2)
	return head
}

func merge(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	dummy := &ListNode{Val: 0}
	p := dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			p.Next = head1
			head1 = head1.Next
		} else {
			p.Next = head2
			head2 = head2.Next
		}

		p = p.Next
	}

	if head1 != nil {
		p.Next = head1
	}

	if head2 != nil {
		p.Next = head2
	}

	return dummy.Next
}
