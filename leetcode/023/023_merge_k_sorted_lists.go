/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午2:51'
*/

package _23

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (h minHeap) Less(i, j int) bool {
	if h[i] == nil {
		return false
	}

	if h[j] == nil {
		return true
	}

	return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	l := len(old)
	r := old[l-1]
	*h = old[:l-1]
	return r
}

func mergeKLists(lists []*ListNode) *ListNode {
	tmp := minHeap(lists)
	h := &tmp
	heap.Init(h)

	var head, last *ListNode

	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		if v == nil {
			continue
		}

		if last != nil {
			last.Next = v
		}

		if head == nil {
			head = v
		}

		last = v
		if v.Next != nil {
			heap.Push(h, v.Next)
		}
	}

	return head
}
