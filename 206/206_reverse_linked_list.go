package _206

import (
	"go_projects/leetcode_golang_version"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午9:35'
*/

type ListNode = leetcode_golang_version.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
