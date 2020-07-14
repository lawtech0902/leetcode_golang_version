/*
__author__ = 'lawtech'
__date__ = '2020/07/13 5:32 下午'
*/

package _693

// 将n右移一位与n异或，检查结果是否全为1
func hasAlternatingBits(n int) bool {
	tmp := n ^ (n >> 1)
	return tmp&(tmp+1) == 0
}
