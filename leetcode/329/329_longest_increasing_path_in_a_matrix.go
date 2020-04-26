/*
__author__ = 'lawtech'
__date__ = '2020/04/26 10:55 上午'
*/

package _329

import "math"

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	walkedLength := make([][]int, len(matrix))
	for i := range walkedLength {
		walkedLength[i] = make([]int, len(matrix[0]))
	}

	res := 0
	for i := range matrix {
		for j := range matrix[0] {
			if m := walk(math.MaxInt32+1, matrix, i, j, walkedLength); m > res {
				res = m
			}
		}
	}

	return res
}

func walk(prev int64, matrix [][]int, i, j int, walkedLength [][]int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return 0
	}

	if prev <= int64(matrix[i][j]) {
		return 0
	}

	if walkedLength[i][j] != 0 {
		return walkedLength[i][j]
	}

	left := walk(int64(matrix[i][j]), matrix, i-1, j, walkedLength) + 1
	right := walk(int64(matrix[i][j]), matrix, i+1, j, walkedLength) + 1
	up := walk(int64(matrix[i][j]), matrix, i, j-1, walkedLength) + 1
	down := walk(int64(matrix[i][j]), matrix, i, j+1, walkedLength) + 1
	walkedLength[i][j] = max(left, right, up, down)

	return walkedLength[i][j]
}

func max(nums ...int) int {
	res := math.MinInt32
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}
