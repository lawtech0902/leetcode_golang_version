/*
__author__ = 'lawtech'
__date__ = '2020/07/30 9:48 上午'
*/

package _24

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func reverseList1(head *ListNode) *ListNode {
	var cur *ListNode
	pre := head
	for pre != nil {
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}

	return cur
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
