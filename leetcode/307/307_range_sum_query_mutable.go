/*
__author__ = 'lawtech'
__date__ = '2020/04/21 2:00 下午'
*/

package _307

type NumArray struct {
	items []int
}

func Constructor(nums []int) NumArray {
	return NumArray{items: nums}
}

func (this *NumArray) Update(i int, val int) {
	this.items[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	subArr := this.items[i : j+1]
	sum := 0

	for _, item := range subArr {
		sum += item
	}

	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
