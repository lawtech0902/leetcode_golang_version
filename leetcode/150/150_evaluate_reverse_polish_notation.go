/*
__author__ = 'lawtech'
__date__ = '2020/03/11 11:18 上午'
*/

package _150

import (
	"container/list"
	"strconv"
)

/*
用栈来存储数字
遍历tokens
入到数字就入栈
遇到一个运算符，弹出栈中的两个元素进行对应计算，然后再把结果压入栈中
*/
func evalRPN(tokens []string) int {
	stack := list.New()
	for _, s := range tokens {
		switch s {
		case "+":
			a, b := popTwoNums(stack)
			stack.PushFront(a + b)
		case "-":
			a, b := popTwoNums(stack)
			stack.PushFront(b - a)
		case "*":
			a, b := popTwoNums(stack)
			stack.PushFront(a * b)
		case "/":
			a, b := popTwoNums(stack)
			stack.PushFront(b / a)
		default:
			num, _ := strconv.Atoi(s)
			stack.PushFront(num)
		}
	}

	return stack.Remove(stack.Front()).(int)
}

func popTwoNums(stack *list.List) (a, b int) {
	a = stack.Remove(stack.Front()).(int)
	b = stack.Remove(stack.Front()).(int)
	return
}
