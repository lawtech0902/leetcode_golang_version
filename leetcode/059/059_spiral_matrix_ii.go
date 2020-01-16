/*
__author__ = 'lawtech'
__date__ = '2020/1/16 11:16 上午'
*/

package _59

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	up, left, down, right := 0, 0, n-1, n-1
	direct, count := 0, 0 // direction: 0: go right   1: go down  2: go left  3: go up

	for {
		switch direct {
		case 0:
			for i := left; i < right+1; i++ {
				count += 1
				matrix[up][i] = count
			}
			up += 1
		case 1:
			for i := up; i < down+1; i++ {
				count += 1
				matrix[i][right] = count
			}
			right -= 1
		case 2:
			for i := right; i > left-1; i-- {
				count += 1
				matrix[down][i] = count
			}
			down -= 1
		case 3:
			for i := down; i > up-1; i-- {
				count += 1
				matrix[i][left] = count
			}
			left += 1
		}

		if count == n*n {
			return matrix
		}

		direct = (direct + 1) % 4
	}
}
