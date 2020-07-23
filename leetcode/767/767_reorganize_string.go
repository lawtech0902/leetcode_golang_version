/*
__author__ = 'lawtech'
__date__ = '2020/07/23 10:16 上午'
*/

package _767

import (
	"container/heap"
	"sort"
)

// 根据出现次数排序
func reorganizeString(S string) string {
	n := len(S)
	counts := make([]int, 26)
	for _, c := range S {
		counts[c-'a'] += 100
	}

	// encode: counts[i] = 100 * (actual count) + i
	for i := 0; i < 26; i++ {
		counts[i] += i
	}

	sort.Ints(counts)
	res := make([]byte, 0)
	t := 1
	for _, code := range counts {
		ct := code / 100
		ch := 'a' + code%100
		if ct > (n+1)/2 {
			return ""
		}

		for i := 0; i < ct; i++ {
			if t >= n {
				t = 0
			}

			res[t] = byte(ch)
			t += 2
		}
	}

	return string(res)
}

// 贪心堆
type rCount struct {
	val   rune
	count int
}

// rune count list
type rList []rCount

func (r rList) Len() int {
	return len(r)
}

func (r rList) Less(i, j int) bool {
	return r[i].count > r[j].count
}

func (r rList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *rList) Push(x interface{}) {
	*r = append(*r, x.(rCount))
}

func (r *rList) Pop() interface{} {
	top := (*r)[0]
	(*r) = (*r)[1:]
	return top
}

func reorganizeString1(S string) string {
	if len(S) <= 1 {
		return S
	}

	// count rune
	count := make(map[rune]int)
	for _, r := range S {
		count[r]++
		if count[r] > (len(S)+1)/2 {
			return ""
		}
	}

	// map to slice
	var nums rList
	for k, v := range count {
		nums = append(nums, rCount{
			val:   k,
			count: v,
		})
	}

	heap.Init(&nums)
	var (
		res  []rune
		prev rune
	)

	for len(nums) > 0 {
		n1 := nums.Pop().(rCount)
		if n1.val != prev {
			res = append(res, n1.val)
			n1.count--
			prev = n1.val
		}

		if len(nums) == 0 {
			break
		}

		n2 := nums.Pop().(rCount)
		if n2.val != prev {
			res = append(res, n2.val)
			n2.count--
			prev = n2.val
		}

		if n1.count > 0 {
			heap.Push(&nums, n1)
		}

		if n2.count > 0 {
			heap.Push(&nums, n2)
		}
	}

	return string(res)
}
