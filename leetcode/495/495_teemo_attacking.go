/*
__author__ = 'lawtech'
__date__ = '2020/06/02 2:05 下午'
*/

package _495

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	if len(timeSeries) == 1 {
		return duration
	}

	res := 0
	for i := 1; i < len(timeSeries); i++ {
		if gap := timeSeries[i] - timeSeries[i-1]; gap < duration {
			res += gap
		} else {
			res += duration
		}
	}

	res += duration
	return res
}
