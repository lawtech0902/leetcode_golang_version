/*
__author__ = 'lawtech'
__date__ = '2020/07/10 10:03 上午'
*/

package _682

import "strconv"

func calPoints(ops []string) int {
	var res []int
	sum := 0
	for _, v := range ops {
		switch v {
		case "+":
			res = append(res, res[len(res)-1]+res[len(res)-2])
		case "C":
			res = res[:len(res)-1]
		case "D":
			res = append(res, res[len(res)-1]*2)
		default:
			point, _ := strconv.Atoi(v)
			res = append(res, point)
		}
	}

	for _, v := range res {
		sum += v
	}

	return sum
}
