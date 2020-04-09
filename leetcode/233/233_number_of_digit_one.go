/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-04-08 15:37:36
 */

package _233

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	m := make(map[int]int)
	m[10] = 2
	for i := 1; i <= 9; i++ {
		m[i] = 1
	}

	return helper(n, m)
}

func helper(n int, m map[int]int) int {
	if n <= 0 {
		return 0
	}

	if ret, ok := m[n]; ok {
		return ret
	}

	ret := 0
	d := highest(n)
	h := n / d // first digit   12345 -> 1
	r := n % d
	if h == 1 {
		ret += r + 1 // highest digit count is 2345 + 1
		ret += helper(r, m)
		ret += helper(d-1, m)
	} else {
		ret += d
		ret += helper(r, m)
		ret += h * helper(d-1, m)
	}

	m[n] = ret
	return ret
}

// for n = 4, return 1
// for n = 33, return 10
// for n = 555, return 100
func highest(n int) int {
	h := 1
	for n/(h*10) > 0 {
		h = h * 10
	}

	return h
}
