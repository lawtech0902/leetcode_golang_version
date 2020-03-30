/*
__author__ = 'lawtech'
__date__ = '2020/03/30 10:15 上午'
*/

package _215

import "container/heap"

// 快排
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	pivot := nums[0]
	p1, p2 := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= pivot {
			p2++
		} else {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
			p2++
		}
	}

	nums[0], nums[p1-1] = nums[p1-1], nums[0]
	if p1 == len(nums)-k+1 {
		return pivot
	} else if p1 < len(nums)-k+1 {
		return findKthLargest(nums[p1:], k)
	} else {
		return findKthLargest(nums[:p1-1], p1+k-len(nums)-1)
	}
}

// 堆
func findKthLargest1(nums []int, k int) int {
	ih := &IntHeap{}
	heap.Init(ih)

	for i := 0; i < k; i++ {
		heap.Push(ih, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < (*ih)[0] {
			continue
		}

		heap.Pop(ih)
		heap.Push(ih, nums[i])
	}

	return (*ih)[0]
}

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[:n-1]
	return x
}
