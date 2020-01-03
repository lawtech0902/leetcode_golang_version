package _24

import "go_projects/leetcode_golang_version"

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午8:22'
*/

type ListNode = leetcode_golang_version.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// newHead指向head.Next节点
	newHead := head.Next
	// 让head.Next指向转换好的newHead.Next节点
	head.Next = swapPairs(newHead.Next)
	// 让newHead.Next指向head节点
	newHead.Next = head

	return newHead
}
