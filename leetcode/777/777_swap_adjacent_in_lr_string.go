/*
__author__ = 'lawtech'
__date__ = '2020/07/23 1:36 下午'
*/

package _777

func canTransform(start, end string) bool {
	curL, curR := 0, 0
	for i := 0; i < len(start); i++ {
		switch start[i] {
		case 'L':
			if curR != 0 {
				return false
			}

			curL--
		case 'R':
			if curL != 0 {
				return false
			}

			curR++
		}

		switch end[i] {
		case 'L':
			curL++
		case 'R':
			curR--
		}

		if curL < 0 || curR < 0 {
			return false
		}

		if curL > 0 && curR > 0 {
			return false
		}
	}

	return curL == 0 && curR == 0
}
