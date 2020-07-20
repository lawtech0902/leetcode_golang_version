/*
__author__ = 'lawtech'
__date__ = '2020/07/20 11:09 上午'
*/

package _735

// stack
func asteroidCollision(asteroids []int) []int {
	res := make([]int, 0)
	for _, v := range asteroids {
		if len(res) == 0 {
			res = append(res, v)
		} else {
			for len(res) > 0 && (res[len(res)-1] > 0 && v < 0) {
				top := res[len(res)-1]
				res = res[:len(res)-1]
				collision := top + v
				if collision > 0 {
					res = append(res, top)
					v = 0
					break
				} else if collision == 0 {
					v = 0
					break
				}
			}

			if v != 0 {
				res = append(res, v)
			}
		}
	}

	return res
}
