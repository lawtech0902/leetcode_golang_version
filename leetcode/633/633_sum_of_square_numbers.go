/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-01-03 16:19:42
 */
package _633

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/24 下午6:03'
*/

// 双指针
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))

	for i <= j {
		powSum := i*i + j*j
		if powSum == c {
			return true
		} else if powSum > c {
			j--
		} else {
			i++
		}
	}

	return false
}

// 费马平方和定理：一个非负整数 cc 能够表示为两个整数的平方和，当且仅当 cc 的所有形如 4k+34k+3 的质因子的幂次均为偶数。
func judgeSquareSum1(c int) bool {
	for i := 2; i*i <= c; i++ {
		count := 0
		if c%i == 0 {
			for c%i == 0 {
				count++
				c /= i
			}

			if i%4 == 3 && count%2 != 0 {
				return false
			}
		}
	}

	return c%4 != 3
}
