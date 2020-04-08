/*
__author__ = 'lawtech'
__date__ = '2018/8/19 上午1:39'
*/

package _232

// two stack
//type MyQueue struct {
//	inStack  []int
//	outStack []int
//}
//
///** Initialize your data structure here. */
//func Constructor() MyQueue {
//	return MyQueue{
//		inStack:  []int{},
//		outStack: []int{},
//	}
//}
//
///** Push element x to the back of queue. */
//func (this *MyQueue) Push(x int) {
//	this.inStack = append(this.inStack, x)
//}
//
///** Removes the element from in front of queue and returns that element. */
//func (this *MyQueue) Pop() int {
//	this.Move()
//	popVal := this.outStack[len(this.outStack)-1]
//	this.outStack = this.outStack[:len(this.outStack)-1]
//	return popVal
//}
//
///** Get the front element. */
//func (this *MyQueue) Peek() int {
//	this.Move()
//	return this.outStack[len(this.outStack)-1]
//}
//
///** Returns whether the queue is empty. */
//func (this *MyQueue) Empty() bool {
//	return len(this.inStack) == 0 && len(this.outStack) == 0
//}
//
///** move the element from inStack to outStack */
//func (this *MyQueue) Move() {
//	if len(this.outStack) == 0 {
//		for len(this.inStack) > 0 {
//			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
//			this.inStack = this.inStack[:len(this.inStack)-1]
//		}
//	}
//}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// one slice
type MyQueue struct {
	arr []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}

	top := this.Peek()
	this.arr = this.arr[1:len(this.arr)]
	return top
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}

	return this.arr[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}
