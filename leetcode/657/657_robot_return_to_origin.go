/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-07 09:37:07
 */

package _657

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, move := range moves {
		if move == 'U' {
			y--
		} else if move == 'D' {
			y++
		} else if move == 'L' {
			x--
		} else {
			x++
		}
	}

	return x == 0 && y == 0
}
