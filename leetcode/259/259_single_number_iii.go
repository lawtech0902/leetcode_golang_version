/*
__author__ = 'lawtech'
__date__ = '2020/04/13 4:49 下午'
*/

package _259

// HashMap
func singleNumber(nums []int) []int {
	mapping := make(map[int]int)
	for _, num := range nums {
		if _, ok := mapping[num]; !ok {
			mapping[num] = 1
		} else {
			mapping[num] += 1
		}
	}

	var res []int
	for num, count := range mapping {
		if count == 1 {
			res = append(res, num)
		}
	}

	return res
}

// bit mask
func singleNumber1(nums []int) []int {
	bitmask := 0
	for _, num := range nums {
		bitmask ^= num
	}

	diff := bitmask & (-bitmask)
	x := 0
	for _, num := range nums {
		if num&diff != 0 {
			x ^= num
		}
	}

	return []int{x, bitmask ^ x}
}
