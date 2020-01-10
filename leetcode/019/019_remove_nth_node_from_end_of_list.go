/*
__author__ = 'lawtech'
__date__ = '2018/8/18 上午1:18'
*/

package _19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 快慢指针
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

	// 链表长度小于n
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
