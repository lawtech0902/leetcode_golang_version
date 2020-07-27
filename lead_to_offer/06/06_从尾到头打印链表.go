/*
__author__ = 'lawtech'
__date__ = '2020/07/27 1:38 下午'
*/

package _6

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	res := reversePrint(head.Next)
	res = append(res, head.Val)
	return res
}

// 翻转结果数组
func reversePrint1(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
