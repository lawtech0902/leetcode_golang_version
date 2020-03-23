/*
__author__ = 'lawtech'
__date__ = '2020/03/23 1:27 下午'
*/

package _202

func isHappy(n int) bool {
	// 快慢指针破循环
	slow, fast := n, n
	for {
		slow = squareSum(slow)
		fast = squareSum(fast)
		fast = squareSum(fast)
		if fast == slow {
			break
		}
	}

	return slow == 1
}

func squareSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}
