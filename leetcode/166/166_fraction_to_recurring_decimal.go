/*
__author__ = 'lawtech'
__date__ = '2020/03/16 9:17 下午'
*/

package _166

import "strconv"

// 恶心人
func fractionToDecimal(numerator int, denominator int) string {
	sign := "" // 符号部分
	if numerator != 0 && (numerator&0x80000000)^(denominator&0x80000000) > 0 {
		sign = "-"
	}
	numerator, denominator = Abs(numerator), Abs(denominator)

	integer := strconv.Itoa(numerator / denominator) // 整数部分
	numerator = numerator % denominator * 10
	if numerator == 0 { // 只有整数部分
		return sign + integer
	}

	decimal := ""         // 小数部分
	m := map[string]int{} // 如果小数点后第 1 位计算了 20/3，则记录："20/3": 0
	for i := 0; numerator > 0; i++ {
		key := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
		if start, ok := m[key]; ok { // 进入了第二次循环，则 [start,i) 就是循环部分
			decimal = decimal[:start] + "(" + decimal[start:i] + ")"
			break
		}
		m[key] = i
		decimal += strconv.Itoa(numerator / denominator)
		numerator = numerator % denominator * 10
	}

	return sign + integer + "." + decimal
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
