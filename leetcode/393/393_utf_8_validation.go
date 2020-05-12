/*
__author__ = 'lawtech'
__date__ = '2020/05/12 1:25 下午'
*/

package _393

func validUtf8(data []int) bool {
	i, l := 0, 0
	n := len(data)
	for i < n {
		if (data[i] & 0x80) == 0 {
			//1 byte: 0xxxxx
			l = 1
		} else if (data[i] & 0xe0) == 0xc0 {
			//2 bytes
			l = 2
		} else if (data[i] & 0xf0) == 0xe0 {
			//3 bytes
			l = 3
		} else if (data[i] & 0xf8) == 0xf0 {
			//4 bytes
			l = 4
		} else {
			return false
		}
		for j := 1; j < l; j++ {
			if (i+j > n-1) || ((data[i+j] & 0xc0) != 0x80) {
				return false
			}
		}

		i += l
	}

	return true
}
