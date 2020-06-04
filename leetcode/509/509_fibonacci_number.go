/*
__author__ = 'lawtech'
__date__ = '2020/06/04 11:20 上午'
*/

package _509

// 1.递归
func fib(N int) int {
	if N <= 1 {
		return N
	}

	return fib(N-1) + fib(N-2)
}

// 2.记忆化自底向上
//func fib2(N int) int {
//	if N <= 1 {
//		return N
//	}
//
//	return memorize(N)
//}
//
//func memorize(N int) int {
//	cache := map[int]int{0: 0, 1: 1}
//	for i := 2; i <= N; i++ {
//		cache[i] = cache[i-1] + cache[i-2]
//	}
//
//	return cache[N]
//}

// 3.记忆化自顶向下
var cache = map[int]int{0: 0, 1: 1}

func fib3(N int) int {
	if N <= 1 {
		return N
	}

	return memorize(N)
}

func memorize(N int) int {
	if _, ok := cache[N]; ok {
		return cache[N]
	}

	cache[N] = memorize(N-1) + memorize(N-2)
	return memorize(N)
}

// 4.自底向上迭代
func fib4(N int) int {
	if N <= 1 {
		return N
	}

	if N == 2 {
		return 1
	}

	cur := 0
	prev1 := 1
	prev2 := 1

	for i := 3; i <= N; i++ {
		cur = prev1 + prev2
		prev2 = prev1
		prev1 = cur
	}

	return cur
}
