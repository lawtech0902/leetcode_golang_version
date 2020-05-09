/*
__author__ = 'lawtech'
__date__ = '2020/05/09 3:37 下午'
*/

package _384

import "math/rand"

type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	return Solution{arr: nums}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.arr
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	tmp := make([]int, len(s.arr))
	copy(tmp, s.arr)
	for i := 0; i < len(s.arr); i++ {
		r := rand.Intn(len(s.arr))
		tmp[i], tmp[r] = tmp[r], tmp[i]
	}

	return tmp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
