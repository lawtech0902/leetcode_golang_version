/*
__author__ = 'lawtech'
__date__ = '2020/07/17 11:07 上午'
*/

package _725

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建新链表
func splitListToParts(root *ListNode, k int) []*ListNode {
	cur := root
	n := 0
	for cur != nil {
		cur = cur.Next
		n++
	}

	width, rem, cnt := n/k, n%k, 0
	res := make([]*ListNode, k)
	cur = root
	for i := 0; i < k; i++ {
		head := &ListNode{}
		write := head
		if i < rem {
			cnt = width + 1
		} else {
			cnt = width
		}
		for j := 0; j < cnt; j++ {
			write.Next = &ListNode{Val: cur.Val}
			write = write.Next
			if cur != nil {
				cur = cur.Next
			}
		}

		res[i] = head.Next
	}

	return res
}

// 拆分链表
func splitListToParts1(root *ListNode, k int) []*ListNode {
	cur := root
	n := 0
	for cur != nil {
		cur = cur.Next
		n++
	}

	width, rem, cnt := n/k, n%k, 0
	res := make([]*ListNode, k)
	cur = root
	for i := 0; i < k; i++ {
		head := cur
		if i < rem {
			cnt = width
		} else {
			cnt = width - 1
		}
		for j := 0; j < cnt; j++ {
			if cur != nil {
				cur = cur.Next
			}
		}

		if cur != nil {
			prev := cur
			cur = cur.Next
			prev.Next = nil
		}

		res[i] = head
	}

	return res
}
