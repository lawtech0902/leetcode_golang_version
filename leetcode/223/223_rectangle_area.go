/*
__author__ = 'lawtech'
__date__ = '2020/04/01 2:12 下午'
*/

package _223

import "math"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1 := (C - A) * (D - B)
	area2 := (G - E) * (H - F)
	xOverlap := overlap(A, C, E, G)
	yOverlap := overlap(B, D, F, H)
	return area1 + area2 - (xOverlap * yOverlap)
}

func overlap(s1, e1, s2, e2 int) int {
	higherStartPoint := int(math.Max(float64(s1), float64(s2)))
	lowerEndPoint := int(math.Min(float64(e1), float64(e2)))

	if lowerEndPoint <= higherStartPoint {
		return 0
	}

	return lowerEndPoint - higherStartPoint
}
