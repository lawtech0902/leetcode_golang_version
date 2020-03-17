/*
__author__ = 'lawtech'
__date__ = '2020/03/16 9:25 下午'
*/

package _168

func convertToTitle(n int) string {
	// 十进制转十六进制
	res := ""
	for n > 0 {
		m := n % 26
		if m == 0 {
			m = 26
			n -= 1
		}
		n = n / 26
		res = string(m+64) + res
	}

	return res
}
