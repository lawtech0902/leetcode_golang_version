/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午11:19'
*/

package _263

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	primeFactors := []int{2, 3, 5}
	for _, x := range primeFactors {
		for num%x == 0 {
			num /= x
		}
	}

	return num == 1
}
