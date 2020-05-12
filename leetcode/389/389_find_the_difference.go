/*
__author__ = 'lawtech'
__date__ = '2020/05/12 10:43 上午'
*/

package _389

// 数字与自身异或结果为0
func findTheDifference(s string, t string) byte {
	var res byte

	for _, c := range []byte(s) {
		res ^= c
	}

	for _, c := range []byte(t) {
		res ^= c
	}

	return res
}
