package _142

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 上午2:08'
*/

type ListNode = leetcode_golang_version.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if slow == fast {
		slow = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}

	return nil
}
