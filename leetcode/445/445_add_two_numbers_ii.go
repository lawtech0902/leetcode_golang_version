/*
__author__ = 'lawtech'
__date__ = '2020/05/21 9:41 上午'
*/

package _445

type ListNode struct {
	Val  int
	Next *ListNode
}

// 不能原地反转链表，就借助栈来反转
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		s1, s2           []int
		a, b, carry, cur int
		res              *ListNode
	)

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		if len(s1) == 0 {
			a = 0
		} else {
			a = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) == 0 {
			b = 0
		} else {
			b = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		cur = a + b + carry
		carry = cur / 10
		cur %= 10
		curNode := &ListNode{Val: cur}
		curNode.Next = res
		res = curNode
	}

	return res
}
