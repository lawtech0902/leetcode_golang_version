/*
__author__ = 'lawtech'
__date__ = '2020/03/09 2:31 下午'
*/

package _143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// split
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head1, head2 := head, slow.Next
	slow.Next = nil

	// reverse
	cur := head2
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	// merge
	cur1, cur2 := head1, pre
	for cur2 != nil {
		tmp1, tmp2 := cur1.Next, cur2.Next
		cur1.Next = cur2
		cur2.Next = tmp1
		cur1, cur2 = tmp1, tmp2
	}
}
