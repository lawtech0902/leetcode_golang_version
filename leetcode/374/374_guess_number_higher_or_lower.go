/*
__author__ = 'lawtech'
__date__ = '2020/05/08 10:55 上午'
*/

package _374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
