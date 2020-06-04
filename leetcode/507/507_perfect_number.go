/*
__author__ = 'lawtech'
__date__ = '2020/06/04 10:18 上午'
*/

package _507

func checkPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}

	sums := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sums += i
			if i*i != num {
				sums += num / i
			}
		}
	}

	return sums-num == num
}
