/*
__author__ = 'lawtech'
__date__ = '2020/07/31 9:50 上午'
*/

package _31

import "container/list"

func validateStackSequences(pushed, popped []int) bool {
	stack := list.New()
	i := 0
	for _, num := range pushed {
		stack.PushBack(num)
		for stack.Len() != 0 && stack.Back().Value == popped[i] {
			stack.Remove(stack.Back())
			i++
		}
	}

	return stack.Len() == 0
}
