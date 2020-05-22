/*
__author__ = 'lawtech'
__date__ = '2020/05/22 10:49 上午'
*/

package _454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	hash := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			hash[a+b]++
		}
	}

	var res int
	for _, c := range C {
		for _, d := range D {
			target := -(c + d)
			if val, ok := hash[target]; ok {
				res += val
			}
		}
	}

	return res
}
