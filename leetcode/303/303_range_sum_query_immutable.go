/*
__author__ = 'lawtech'
__date__ = '2020/04/21 10:30 上午'
*/

package _303

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	this := NumArray{
		nums: make([]int, len(nums)+1),
	}

	for i := 1; i <= len(nums); i++ {
		this.nums[i] = this.nums[i-1] + nums[i-1]
	}

	return this
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.nums[j+1] - this.nums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
