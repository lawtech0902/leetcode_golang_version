/*
__author__ = 'lawtech'
__date__ = '2020/07/13 1:48 下午'
*/

package _692

import (
	"container/heap"
	"sort"
	"strings"
)

// 排序
func topKFrequent(words []string, k int) []string {
	hash := make(map[string]int)
	for _, word := range words {
		hash[word]++
	}

	counter := make([][2]interface{}, 0)
	for k, v := range hash {
		counter = append(counter, [2]interface{}{k, v})
	}

	sort.Slice(counter, func(i, j int) bool {
		if counter[i][1] == counter[j][1] {
			c := strings.Compare(counter[i][0].(string), counter[j][0].(string))
			if c <= 0 {
				return true
			}

			return false
		}

		return counter[i][1].(int) > counter[j][1].(int)
	})

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = counter[i][0].(string)
	}

	return res
}

// 堆
func topKFrequent1(words []string, k int) []string {
	hash := make(map[string]int)
	for _, word := range words {
		hash[word]++
	}

	var mh MinHeap
	heap.Init(&mh)
	for key := range hash {
		heap.Push(&mh, &Word{key, hash[key]})
		if mh.Len() > k {
			heap.Pop(&mh)
		}
	}

	res := make([]string, k, k)
	for mh.Len() > 0 {
		res[mh.Len()-1] = heap.Pop(&mh).(*Word).val
	}

	return res
}

type Word struct {
	val string
	cnt int
}

type MinHeap []*Word

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].cnt == h[j].cnt {
		return h[i].val > h[j].val
	}

	return h[i].cnt < h[j].cnt
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Word))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
