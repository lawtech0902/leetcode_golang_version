/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午11:32'
*/

package _160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(pHead1, pHead2 *ListNode) *ListNode {
	if pHead1 != nil && pHead2 != nil {
		p1, p2 := pHead1, pHead2
		for p1 != p2 {
			if p1 != nil {
				p1 = p1.Next
			} else {
				p1 = pHead2
			}

			if p2 != nil {
				p2 = p2.Next
			} else {
				p2 = pHead1
			}
		}

		return p1
	}

	return nil
}
