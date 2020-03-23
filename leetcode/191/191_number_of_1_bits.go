/*
__author__ = 'lawtech'
__date__ = '2020/03/19 10:36 上午'
*/

package _191

// 位掩码
func hammingWeight1(num uint32) int {
	res, mask := 0, 1
	for i := 0; i < 32; i++ {
		if num&uint32(mask) != 0 {
			res++
		}

		mask <<= 1
	}

	return res
}

// 位反转
func hammingWeight2(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num &= num - 1
	}

	return res
}
