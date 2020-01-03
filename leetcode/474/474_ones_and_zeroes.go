package _474

import "math"

/*
__author__ = 'lawtech'
__date__ = '2018/8/21 上午11:05'
*/

// dp[x][y]表示至多使用x个0，y个1可以组成字符串的最大数目
func findMaxForm(strs []string, m int, n int) int {
	if strs == nil || len(strs) == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zero, one := getNumber(s)
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = int(math.Max(float64(dp[i-zero][j-one]+1), float64(dp[i][j])))
			}
		}
	}

	return dp[m][n]
}

func getNumber(str string) (int, int) {
	zero, one := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}
