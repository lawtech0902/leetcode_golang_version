/*
__author__ = 'lawtech'
__date__ = '2020/07/27 3:03 下午'
*/

package _9

import "container/list"

type CQueue struct {
	stackIn  *list.List
	stackOut *list.List
}

func Constructor() CQueue {
	return CQueue{
		stackIn:  list.New(),
		stackOut: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stackOut.Len() == 0 {
		for this.stackIn.Len() > 0 {
			this.stackOut.PushBack(this.stackIn.Remove(this.stackIn.Back()))
		}
	}

	if this.stackOut.Len() != 0 {
		e := this.stackOut.Back()
		this.stackOut.Remove(e)
		return e.Value.(int)
	}

	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
