package _21

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午1:24'
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head, pHead *ListNode

	if l1.Val < l2.Val {
		head = l1
		pHead = l1
		l1 = l1.Next
	} else {
		head = l2
		pHead = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pHead.Next = l1
			l1 = l1.Next
		} else {
			pHead.Next = l2
			l2 = l2.Next
		}

		pHead = pHead.Next
	}

	if l1 != nil {
		pHead.Next = l1
	}

	if l2 != nil {
		pHead.Next = l2
	}

	return head
}
