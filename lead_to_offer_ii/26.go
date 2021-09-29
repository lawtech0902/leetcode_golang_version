package lead_to_offer_ii

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 快慢指针找中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head1 := slow.Next
	slow.Next = nil

	// 反转后半段
	cur := head1
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	// 归并
	cur1, cur2 := head, pre
	for cur2 != nil {
		temp1, temp2 := cur1.Next, cur2.Next
		cur1.Next = cur2
		cur2.Next = temp1
		cur1, cur2 = temp1, temp2
	}
}
