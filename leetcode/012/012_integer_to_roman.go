/*
__author__ = 'lawtech'
__date__ = '2020/1/9 4:54 下午'
*/

package _12

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""

	for i := range values {
		for num >= values[i] {
			num -= values[i]
			res += numerals[i]
		}
	}

	return res
}
