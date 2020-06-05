/*
__author__ = 'lawtech'
__date__ = '2020/06/05 1:32 下午'
*/

package _521

import "math"

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	return int(math.Max(float64(len(a)), float64(len(b))))
}
