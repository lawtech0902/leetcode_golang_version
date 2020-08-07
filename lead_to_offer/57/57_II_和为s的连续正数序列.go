/*
__author__ = 'lawtech'
__date__ = '2020/08/07 1:49 下午'
*/

package _57

// 滑动窗口
func findContinuousSequence(target int) [][]int {
	i, j := 1, 1 // 左右边界
	sum := 0
	var res [][]int

	for i <= target/2 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			var arr []int
			for k := i; k < j; k++ {
				arr = append(arr, k)
			}

			res = append(res, arr)
			sum -= i
			i++
		}
	}

	return res
}
