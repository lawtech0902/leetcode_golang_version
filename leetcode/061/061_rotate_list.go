/*
__author__ = 'lawtech'
__date__ = '2020/1/16 2:05 下午'
*/

package _61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	p := dummy

	size := 0
	for p.Next != nil {
		p = p.Next
		size += 1
	}

	p.Next = dummy.Next
	step := size - (k % size)
	for i := 0; i < step; i++ {
		p = p.Next
	}

	head = p.Next
	p.Next = nil
	return head
}
