/*
__author__ = 'lawtech'
__date__ = '2020/07/13 9:42 上午'
*/

package _688

// dp
func knightProbability(N, K, r, c int) float64 {
	dp := make([][]float64, N)
	for i := range dp {
		dp[i] = make([]float64, N)
	}

	dr := []int{2, 2, 1, 1, -1, -1, -2, -2}
	dc := []int{1, -1, 2, -2, 2, -2, 1, -1}

	dp[r][c] = 1
	for K > 0 {
		dp2 := make([][]float64, N) // dp2 -> f[][][steps] 马在位置 (r, c) 移动了 steps 次以后还留在棋盘上的概率
		for i := range dp2 {
			dp2[i] = make([]float64, N)
		}
		for sr := 0; sr < N; sr++ {
			for sc := 0; sc < N; sc++ {
				for k := 0; k < 8; k++ {
					cr := sr + dr[k]
					cc := sc + dc[k]
					if 0 <= cr && cr < N && 0 <= cc && cc < N {
						dp2[cr][cc] += dp[sr][sc] / 8.0
					}
				}
			}
		}

		dp = dp2
		K--
	}

	var res float64
	for _, row := range dp {
		for _, val := range row {
			res += val
		}
	}

	return res
}

// dfs + memo
type memoKey = [3]int

func knightProbability1(N, K, r, c int) float64 {
	memo := make(map[memoKey]float64)
	moves := [][2]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}}
	var dfs func(K, r, c int) float64

	dfs = func(K, r, c int) float64 {
		if r < 0 || c < 0 || r >= N || c >= N {
			return 0
		}

		if K == 0 {
			return 1
		}

		triple := memoKey{K, r, c}
		if val, ok := memo[triple]; ok {
			return val
		}

		var p float64
		for _, move := range moves {
			p += dfs(K-1, r+move[0], c+move[1])
		}

		p /= 8.0
		memo[triple] = p
		return p
	}

	return dfs(K, r, c)
}
