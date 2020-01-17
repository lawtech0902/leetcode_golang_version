/*
__author__ = 'lawtech'
__date__ = '2020/1/17 2:16 下午'
*/

package _82

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 递归
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}

		return deleteDuplicates(head.Next)
	}

	head.Next = deleteDuplicates(head.Next)
	return head
}
