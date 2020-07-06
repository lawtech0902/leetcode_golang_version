/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-03 11:13:32
 */

package _646

import (
	"math"
	"sort"
)

// dp O(n2) O(n)
func findLongestChain(pairs [][]int) int {
	sort.Sort(helper(pairs))
	res := 0
	rBound := math.MinInt32
	for _, v := range pairs {
		if v[0] > rBound {
			res++
			rBound = v[1]
		}
	}

	return res
}

type helper [][]int

func (h helper) Less(i, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][0] < h[j][0]
	}

	return h[i][1] < h[j][1]
}

func (h helper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h helper) Len() int {
	return len(h)
}
