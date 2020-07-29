/*
__author__ = 'lawtech'
__date__ = '2020/07/29 9:53 上午'
*/

package _15

// 逐位判断
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res += int(num & 1)
		num >>= 1
	}

	return res
}

// 巧用n&(n-1)
func hammingWeight1(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num &= num - 1 // 循环消去最右边的1
	}

	return res
}
