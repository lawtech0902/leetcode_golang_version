/*
__author__ = 'lawtech'
__date__ = '2020/07/30 2:16 下午'
*/

package _29

import "reflect"

// 转圈遍历，效率有点差
func spiralOrder(matrix [][]int) []int {
	var res []int
	for len(matrix) != 0 {
		res = append(res, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) == 0 {
			break
		}

		matrix = turn(matrix)
	}

	return res
}

func turn(matrix [][]int) [][]int {
	rowSize, colSize := len(matrix), len(matrix[0])
	var newMatrix [][]int

	for j := 0; j < colSize; j++ {
		var newSubMatrix []int
		for i := 0; i < rowSize; i++ {
			newSubMatrix = append(newSubMatrix, matrix[i][j])
		}

		newMatrix = append(newMatrix, newSubMatrix)
	}

	reverseAny(newMatrix)
	return newMatrix
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// 模拟、设定边界，大佬牛逼
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// 左右上下四个边界
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	var res []int

	for {
		// 从左到右
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}

		t++
		if t > b {
			break
		}

		// 从上到下
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}

		r--
		if l > r {
			break
		}

		// 从右到左
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}

		b--
		if t > b {
			break
		}

		// 从下到上
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}

		l++
		if l > r {
			break
		}
	}

	return res
}
