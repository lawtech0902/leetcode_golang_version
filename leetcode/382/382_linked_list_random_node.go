/*
__author__ = 'lawtech'
__date__ = '2020/05/09 1:35 下午'
*/

package _382

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
蓄水池抽样算法
每次只保留一个数，当遇到第 i 个数时，以 1/i的概率保留它，(i-1)/i的概率保留原来的数。
*/
type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

/** Returns a random node's value. */
func (s *Solution) GetRandom() int {
	count := 0
	res := 0
	cur := s.head
	for cur != nil {
		count += 1
		r := rand.Intn(count) + 1
		if r == count {
			res = cur.Val
		}

		cur = cur.Next
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
