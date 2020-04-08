/*
__author__ = 'lawtech'
__date__ = '2020/04/08 10:00 上午'
*/

package _225

// 双队列法
type MyStack struct {
	enqueue []int
	dequeue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		enqueue: []int{},
		dequeue: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.enqueue = append(this.enqueue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	size := len(this.enqueue)
	for i := 0; i < size-1; i++ {
		this.dequeue = append(this.dequeue, this.enqueue[0])
		this.enqueue = this.enqueue[1:]
	}

	topEle := this.enqueue[0]
	this.enqueue = this.dequeue
	this.dequeue = nil

	return topEle
}

/** Get the top element. */
func (this *MyStack) Top() int {
	topEle := this.Pop()
	this.enqueue = append(this.enqueue, topEle)
	return topEle
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.enqueue) == 0 {
		return true
	}

	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
