package _26

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午8:36'
*/

func removeDuplicates(a []int) int {
	left, right, size := 0, 1, len(a)

	for ; right < size; right++ {
		if a[left] == a[right] {
			continue
		}

		left ++
		a[left], a[right] = a[right], a[left]
	}

	return left + 1
}
