/*
__author__ = 'lawtech'
__date__ = '2020/07/30 3:13 下午'
*/

package _30

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.Min() < x {
		min = this.Min()
	}

	this.stack = append(this.stack, item{min: min, x: x})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].x
}

func (this *MinStack) Min() int {
	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
