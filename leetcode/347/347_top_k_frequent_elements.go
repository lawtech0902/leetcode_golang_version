/*
__author__ = 'lawtech'
__date__ = '2020/04/29 2:37 下午'
*/

package _347

import "container/heap"

type info struct {
	number int
	freq   int
}

type minHeap []info

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].freq < m[j].freq
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(info))
}

func (m *minHeap) Pop() interface{} {
	var x interface{}
	x, *m = (*m)[len(*m)-1], (*m)[:len(*m)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	cache := make(map[int]int)
	for _, n := range nums {
		cache[n]++
	}

	minHeap := minHeap{}
	heap.Init(&minHeap)
	for num, freq := range cache {
		if len(minHeap) < k {
			heap.Push(&minHeap, info{
				number: num,
				freq:   freq,
			})
		} else if minHeap[0].freq < freq {
			minHeap[0] = info{
				number: num,
				freq:   freq,
			}

			heap.Fix(&minHeap, 0)
		}
	}

	res := make([]int, 0, k)
	for _, info := range minHeap {
		res = append(res, info.number)
	}

	return res
}
