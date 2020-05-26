/*
__author__ = 'lawtech'
__date__ = '2020/05/26 2:38 下午'
*/

package _470

func rand10() int {
	for {
		num := (rand7()-1)*7 + rand7()
		if num <= 40 {
			return 1 + num%10
		}

		num = (num-40-1)*7 + rand7()
		if num <= 60 {
			return 1 + num%10
		}

		num = (num-60-1)*7 + rand7()
		if num <= 20 {
			return 1 + num%10
		}
	}
}
