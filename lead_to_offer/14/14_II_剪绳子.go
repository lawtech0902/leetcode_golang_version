/*
__author__ = 'lawtech'
__date__ = '2020/07/28 5:22 下午'
*/

package _14

// 快速幂
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	b := n % 3
	p := 1000000007
	rem, x := 1, 3
	for a := n/3 - 1; a > 0; a /= 2 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}

		x = (x * x) % p
	}

	if b == 0 {
		return rem * 3 % p
	} else if b == 1 {
		return rem * 4 % p
	} else {
		return rem * 6 % p
	}
}
