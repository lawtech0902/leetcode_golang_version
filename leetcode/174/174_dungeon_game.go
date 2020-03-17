/*
__author__ = 'lawtech'
__date__ = '2020/03/17 10:45 上午'
*/

package _174

/*
dp[i][j] = 从 i，j 到终点需要的最小血量
如果 dp[i][j] <= 0 需要保证 dp[i][j] = 1
*/
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	dp := make([][]int, m)

	if m == 0 {
		return 0
	}

	n := len(dungeon[0])
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if dungeon[m-1][n-1] <= 0 {
		dp[m-1][n-1] = abs(1 - dungeon[m-1][n-1])
	} else {
		dp[m-1][n-1] = 1
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i != m-1 && j != n-1 {
				dp[i][j] = min(dp[i][j+1], dp[i+1][j])
			} else if i == m-1 && j == n-1 {
				continue
			} else if i == m-1 {
				dp[i][j] = dp[i][j+1]
			} else if j == n-1 {
				dp[i][j] = dp[i+1][j]
			}

			dp[i][j] -= dungeon[i][j]
			if dp[i][j] <= 0 {
				dp[i][j] = 1
			}
		}
	}

	return dp[0][0]
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
