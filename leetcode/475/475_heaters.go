/*
__author__ = 'lawtech'
__date__ = '2020/05/27 10:56 上午'
*/

package _475

import "sort"

func findRadius(houses []int, heaters []int) int {
	var distanceArr []int
	sort.Ints(houses)
	sort.Ints(heaters)

	for _, house := range houses {
		// 二分查找，在heaters中寻找与房屋 c 最近的加热器
		left := 0
		right := len(heaters) - 1
		for left < right {
			mid := (left + right) >> 1
			if heaters[mid] < house {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 若找到的值等于 c ，则说明 c 房屋处放有一个加热器，c 房屋到加热器的最短距离为 0
		if heaters[left] == house {
			distanceArr = append(distanceArr, 0)
		} else if heaters[left] < house {
			distanceArr = append(distanceArr, house-heaters[left])
			// 若该加热器的坐标值大于 c 并且left不等于 0 ，说明 c 介于left和left-1之间，
			// 房屋到加热器的最短距离就是left和left - 1处加热器与 c 差值的最小值
		} else if left != 0 {
			distanceArr = append(distanceArr, min(heaters[left]-house, house-heaters[left-1]))
		} else {
			distanceArr = append(distanceArr, heaters[left]-house)
		}
	}

	return maxIntSlice(distanceArr)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxIntSlice(nums []int) int {
	var res int
	for i, num := range nums {
		if i == 0 || num > res {
			res = num
		}
	}

	return res
}
