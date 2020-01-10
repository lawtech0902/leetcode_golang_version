/*
__author__ = 'lawtech'
__date__ = '2020/1/10 2:20 下午'
*/

package _25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := ListNode{Val: 0, Next: head}
	pre := &dummy
	cur := pre
	size := 0

	for cur.Next != nil {
		size += 1
		cur = cur.Next
	}

	for size >= k {
		cur = pre.Next
		for i := 1; i < k; i++ {
			tmp := cur.Next
			cur.Next = tmp.Next
			tmp.Next = pre.Next
			pre.Next = tmp
		}

		pre = cur
		size -= k
	}

	return dummy.Next
}
