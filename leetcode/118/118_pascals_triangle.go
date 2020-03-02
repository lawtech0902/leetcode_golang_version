/*
__author__ = 'lawtech'
__date__ = '2020/02/27 4:05 下午'
*/

package _118

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	// 生成三角形的边
	for i := range res {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		res[i] = tmp
	}

	// 计算三角形中的数字
	for i := 2; i < numRows; i++ {
		// i为行数
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}

	return res
}
