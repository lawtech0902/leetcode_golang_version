package 剑指offer

/*
__author__ = 'lawtech'
__date__ = '2018/8/26 下午9:15'
*/

func fibonacci(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}
