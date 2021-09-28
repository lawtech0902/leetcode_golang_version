package lead_to_offer_ii

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	fast, i := head, 0
	for ; i < n; i++ {
		if fast == nil {
			break
		}

		fast = fast.Next
	}

	if i < n {
		return head
	}

	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
