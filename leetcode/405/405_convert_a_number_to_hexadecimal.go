/*
__author__ = 'lawtech'
__date__ = '2020/05/13 10:36 上午'
*/

package _405

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hash := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	res := ""
	for num != 0 {
		res = string(hash[num&15]) + res
		num = int(uint32(num) >> 4)
	}

	return res
}
