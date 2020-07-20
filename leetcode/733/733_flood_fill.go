/*
__author__ = 'lawtech'
__date__ = '2020/07/20 10:57 上午'
*/

package _733

// dfs
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var dfs func(sr, sc, oldColor int)
	dfs = func(sr, sc, oldColor int) {
		if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldColor {
			return
		}

		image[sr][sc] = newColor
		dfs(sr-1, sc, oldColor)
		dfs(sr+1, sc, oldColor)
		dfs(sr, sc-1, oldColor)
		dfs(sr, sc+1, oldColor)
	}

	if image[sr][sc] != newColor {
		dfs(sr, sc, image[sr][sc])
	}

	return image
}
