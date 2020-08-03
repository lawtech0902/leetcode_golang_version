/*
__author__ = 'lawtech'
__date__ = '2020/08/03 2:02 下午'
*/

package _40

import (
	"container/heap"
	"sort"
)

// 直接sort O(nlogn) O(logn)
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

// 最大堆 O(nlogk) O(k)
func getLeastNumbers1(arr []int, k int) []int {
	var res []int
	if k == 0 || k > len(arr) {
		return res
	}

	h := &intHeap{}
	heap.Init(h)
	for _, v := range arr {
		if h.Len() < k {
			heap.Push(h, v)
		} else {
			if (*h)[0] > v {
				heap.Pop(h)
				heap.Push(h, v)
			}
		}
	}

	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}

	return res
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

// 快排
func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}

	left := 0
	right := len(arr) - 1
	// 1.此处的任务就是去找到基准值
	index := partation(arr, left, right)
	// 如果基准值 不等于K的值
	for index != k-1 {
		// 基准值比 k 大,基准值前面的数组都比K
		if index > k-1 {
			right = index - 1
			index = partation(arr, left, right)
		} else {
			// 基准比K小
			left = index + 1
			index = partation(arr, left, right)
		}
	}

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, arr[i])
	}

	return result

}

func partation(num []int, left, right int) int {
	if left == right {
		return right
	}

	if left < right {
		// 1.找到基准元素
		base := num[left]
		l := left
		r := right
		for {
			// 2. 如果右侧的值大于base ，尾指针 向前移动
			for num[r] >= base && l < r {
				r--
			}
			// 3.如果左侧的值小于base , 头指针向后移动
			for num[l] <= base && l < r {
				l++
			}
			// 交换一下
			if l >= r {
				//fmt.Println("break")
				break

			}

			num[l], num[r] = num[r], num[l]
		}

		// 基准值为找到的那个值
		num[left] = num[l]
		num[l] = base
		return l
	}

	return -1
}
