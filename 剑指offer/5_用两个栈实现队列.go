package 剑指offer

/*
__author__ = 'lawtech'
__date__ = '2018/8/26 下午3:48'
*/

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (mq *MyQueue) Push(x int) {
	mq.inStack = append(mq.inStack, x)
}

func (mq *MyQueue) Peek() int {
	mq.Move()
	return mq.outStack[len(mq.outStack)-1]
}

func (mq *MyQueue) Move() {
	if len(mq.outStack) == 0 {
		for len(mq.inStack) > 0 {
			mq.outStack = append(mq.outStack, mq.inStack[len(mq.inStack)-1])
			mq.inStack = mq.inStack[:len(mq.inStack)-1]
		}
	}
}

func (mq *MyQueue) Empty() bool {
	return len(mq.inStack) == 0 && len(mq.outStack) == 0
}
