/*
__author__ = 'lawtech'
__date__ = '2020/07/29 5:16 下午'
*/

package _22

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
