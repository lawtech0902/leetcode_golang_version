/*
__author__ = 'lawtech'
__date__ = '2020/06/02 9:42 上午'
*/

package _492

import "math"

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w != 0 {
		w -= 1
	}

	return []int{area / w, w}
}
