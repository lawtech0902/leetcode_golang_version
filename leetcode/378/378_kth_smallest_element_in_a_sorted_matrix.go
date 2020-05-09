/*
__author__ = 'lawtech'
__date__ = '2020/05/09 10:08 上午'
*/

package _378

/*
二分查找:将查找上下限设为矩阵的右下角与左上角元素，查找过程中对中值在矩阵每一行的位置进行累加，记该值为loc，根据loc与k的大小关系调整上下限
*/
func kthSmallest(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	low, high := matrix[0][0], matrix[m-1][n-1]
	for low < high {
		mid := low + (high-low)/2
		count := 0
		i, j := 0, n-1
		for ; i < m; i++ {
			for j >= 0 && mid < matrix[i][j] {
				j--
			}

			count += j + 1
		}

		if count < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
