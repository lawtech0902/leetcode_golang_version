/*
__author__ = 'lawtech'
__date__ = '2020/04/21 10:45 上午'
*/

package _304

type NumMatrix struct {
	dp [][]int
}

// sum(abcd) = sum(od) − sum(ob) − sum(oc) + sum(oa)
func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{}
	if len(matrix) == 0 {
		return nm
	}

	nm.dp = make([][]int, len(matrix)+1)
	for i := 0; i < len(nm.dp); i++ {
		nm.dp[i] = make([]int, len(matrix[0])+1)
	}

	for i, mv := range matrix {
		for j, v := range mv {
			nm.dp[i+1][j+1] = v + nm.dp[i+1][j] + nm.dp[i][j+1] - nm.dp[i][j]
		}
	}

	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 < 0 || col1 < 0 {
		return 0
	}

	if row2 >= len(this.dp) || col2 >= len(this.dp[0]) {
		return 0
	}

	return this.dp[row2+1][col2+1] - this.dp[row1][col2+1] - this.dp[row2+1][col1] + this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
