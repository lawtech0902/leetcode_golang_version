package _263

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午11:19'
*/

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for _, x := range []int{2, 3, 5} {
		for num%x == 0 {
			num /= x
		}
	}

	return num == 1
}
