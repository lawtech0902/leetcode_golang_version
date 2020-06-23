/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-23 10:07:35
 */

package _592

import (
	"strconv"
	"strings"
)

func fractionAddition(expression string) string {
	numerator, denominators := []int{}, []int{}
	sign := []int{}
	if expression[0] != '-' {
		sign = append(sign, 1)
	}
	for _, v := range expression {
		if v == '+' {
			sign = append(sign, 1)
		} else if v == '-' {
			sign = append(sign, -1)
		}
	}
	data := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(expression, "+", " "), "-", " "))
	for i, d := range data {
		v := strings.Split(d, "/")
		n, _ := strconv.Atoi(v[0])
		d, _ := strconv.Atoi(v[1])
		numerator = append(numerator, n*sign[i])
		denominators = append(denominators, d)
	}
	resultNumerator := numerator[0]
	resultDenominators := denominators[0]
	for i := 1; i < len(numerator); i++ {
		resultNumerator, resultDenominators = add(resultNumerator, resultDenominators, numerator[i], denominators[i])
	}
	return strconv.Itoa(resultNumerator) + "/" + strconv.Itoa(resultDenominators)
}

func add(numerator1, denominator1, numerator2, denominator2 int) (int, int) {
	if numerator1 == 0 && numerator2 == 0 {
		return 0, 1
	}
	if numerator1 == 0 {
		return numerator2, denominator2
	}
	if numerator2 == 0 {
		return numerator1, denominator1
	}
	denominator := lcm(denominator1, denominator2)
	numerator := denominator/denominator1*numerator1 + denominator/denominator2*numerator2
	if numerator == 0 {
		return 0, 1
	}
	gcd := gcd(numerator, denominator)
	return numerator / gcd, denominator / gcd
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
