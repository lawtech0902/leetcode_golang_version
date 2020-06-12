/*
__author__ = 'lawtech'
__date__ = '2020/06/12 5:13 下午'
*/

package _537

import "fmt"

func complexNumberMultiply(a string, b string) string {
	var p, q, m, n int
	fmt.Sscanf(a, "%d+%di", &p, &q)
	fmt.Sscanf(b, "%d+%di", &m, &n)
	return fmt.Sprintf("%d+%di", p*m-q*n, m*q+p*n)
}
