/*
__author__ = 'lawtech'
__date__ = '2020/07/15 9:59 上午'
*/

package _703

import "container/heap"

type KthLargest struct {
	Nums *IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	ih := &IntHeap{}
	heap.Init(ih)
	for i := 0; i < len(nums); i++ {
		heap.Push(ih, nums[i])
	}

	for len(*ih) > k {
		heap.Pop(ih)
	}

	return KthLargest{
		Nums: ih,
		K:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else if val > (*this.Nums)[0] {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}

	return (*this.Nums)[0]
}

// min heap
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
	tmp := *ih
	n := len(tmp)
	x := tmp[n-1]
	*ih = tmp[:n-1]
	return x
}
