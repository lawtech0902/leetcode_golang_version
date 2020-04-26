/*
__author__ = 'lawtech'
__date__ = '2020/04/26 10:42 上午'
*/

package _328

type ListNode struct {
	Val  int
	Next *ListNode
}

// 画图
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd, even, evenHead := head, head.Next, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
