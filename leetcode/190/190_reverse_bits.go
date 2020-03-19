/*
__author__ = 'lawtech'
__date__ = '2020/03/19 10:30 上午'
*/

package _190

func reverseBits1(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

func reverseBits(num uint32) uint32 {
	// 位左移高效一点
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) + (num & 1)
		num = num >> 1
	}
	return res
}
