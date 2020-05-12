/*
__author__ = 'lawtech'
__date__ = '2020/05/12 10:25 上午'
*/

package _388

import "strings"

func lengthLongestPath(input string) int {
	cur := 0
	levelMap := make(map[int]int)
	levelMap[0] = -1
	levelMap[-1] = -2 // if getLevel return 0,0, levelMap[0] still -1
	paths := strings.Split(input, "\n")
	for _, path := range paths {
		level, length := getLevel(path)
		levelMap[level] = levelMap[level-1] + length + 1
		if isFile(path) {
			if levelMap[level] > cur {
				cur = levelMap[level]
			}
		}
	}

	return cur
}

func getLevel(s string) (int, int) {
	for i, v := range s {
		if v != '\t' {
			return i + 1, len(s) - i
		}
	}

	return 0, 0
}

func isFile(s string) bool {
	return strings.Contains(s, ".")
}
