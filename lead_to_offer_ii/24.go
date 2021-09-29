package lead_to_offer_ii

// 迭代
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

// 递归
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
