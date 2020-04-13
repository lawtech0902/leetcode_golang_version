/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-04-09 16:03:21
 */

package _234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}

	return true
}

func isPalindrome1(head *ListNode) bool {
	fast, slow := head, head
	stack := []int{}

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != stack[len(stack)-1] {
			return false
		}

		slow = slow.Next
		stack = stack[:len(stack)-1]
	}

	return true
}
