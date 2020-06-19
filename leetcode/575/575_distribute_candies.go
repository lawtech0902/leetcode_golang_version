/*
__author__ = 'lawtech'
__date__ = '2020/06/18 11:02 上午'
*/

package _575

func distributeCandies(candies []int) int {
	res := len(candies) / 2
	hash := make(map[int]int)
	for _, candy := range candies {
		hash[candy]++
	}

	kind := len(hash)
	if kind < res {
		res = kind
	}

	return res
}
