/*
__author__ = 'lawtech'
__date__ = '2020/05/12 9:56 上午'
*/

package _386

// 前缀树
func lexicalOrder(n int) []int {
	res := make([]int, 0)
	num := 1

	for {
		if num <= n {
			res = append(res, num)
			num *= 10
		} else {
			num /= 10
			for num%10 == 9 {
				num /= 10
			}

			if num == 0 {
				break
			}

			num++
		}
	}

	return res
}
