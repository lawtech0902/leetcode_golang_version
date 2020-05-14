/*
__author__ = 'lawtech'
__date__ = '2020/05/13 4:31 下午'
*/

package _413

func numberOfArithmeticSlices(A []int) int {
	size := len(A)
	if size < 3 {
		return 0
	}

	res, cnt := 0, 0
	delta := A[1] - A[0]
	for i := 2; i < size; i++ {
		if A[i]-A[i-1] == delta {
			cnt += 1
			res += cnt
		} else {
			delta = A[i] - A[i-1]
			cnt = 0
		}
	}

	return res
}
