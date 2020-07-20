/*
__author__ = 'lawtech'
__date__ = '2020/07/20 1:55 下午'
*/

package _739

// 递减栈
func dailyTemperatures(T []int) []int {
	size := len(T)
	res := make([]int, size)
	stack := make([]int, 0)
	for i := 0; i < size; i++ {
		for len(stack) != 0 && T[i] > T[stack[len(stack)-1]] {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[t] = i - t
		}

		stack = append(stack, i)
	}

	return res
}

// 暴力
func dailyTemperatures1(T []int) []int {
	size := len(T)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		if T[i] < 100 {
			for j := i + 1; j < size; j++ {
				if T[j] > T[i] {
					res[i] = j - i
					break
				}
			}
		}
	}

	return res
}
