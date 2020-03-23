/*
__author__ = 'lawtech'
__date__ = '2020/03/23 2:34 下午'
*/

package _204

func countPrimes(n int) int {
	count := 0
	isPrimes := make([]bool, n)
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			continue
		}

		count++
		for j := 2 * i; j < n; j += i {
			isPrimes[j] = true
		}
	}

	return count
}
