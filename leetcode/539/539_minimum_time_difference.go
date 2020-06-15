/*
__author__ = 'lawtech'
__date__ = '2020/06/15 10:59 上午'
*/

package _539

import (
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	minutes := make([]int, len(timePoints))
	for i, tp := range timePoints {
		minutes[i] = transTimePoint2Minutes(tp)
	}

	sort.Ints(minutes)
	min := minutes[0] + 1440 - minutes[len(minutes)-1]
	for i := 1; i < len(minutes); i++ {
		cur := minutes[i] - minutes[i-1]
		if cur < min {
			min = cur
		}
	}

	return min
}

func transTimePoint2Minutes(timePoint string) int {
	hourMinutes := strings.Split(timePoint, ":")
	h, _ := strconv.Atoi(hourMinutes[0])
	m, _ := strconv.Atoi(hourMinutes[1])
	return h*60 + m
}
