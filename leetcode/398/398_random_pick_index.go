/*
__author__ = 'lawtech'
__date__ = '2020/05/12 4:15 下午'
*/

package _398

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	res := -1
	temp := 1
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			if res == -1 {
				res = i
			} else {
				temp++
				if rand.Intn(temp) == 0 {
					res = i
				}
			}
		}
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
