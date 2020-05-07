/*
__author__ = 'lawtech'
__date__ = '2020/05/07 9:52 上午'
*/

package _365

func canMeasureWater(x int, y int, z int) bool {
	return z == 0 || (x+y >= z && z%(gcd(x, y)) == 0)
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}

	return gcd(y, x%y)
}
