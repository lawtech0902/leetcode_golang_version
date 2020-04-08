/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午10:31'
*/

package _229

func majorityElement(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	e1, e2, c1, c2 := 0, 1, 0, 0

	for _, num := range nums {
		switch {
		case num == e1:
			c1++
		case num == e2:
			c2++
		case c1 == 0:
			e1 = num
			c1 = 1
		case c2 == 0:
			e2 = num
			c2 = 1
		default:
			c1--
			c2--
		}
	}

	var res []int

	if maj(nums, e2) {
		res = append(res, e2)
	}

	if maj(nums, e1) {
		res = append(res, e1)
	}

	return res
}

func maj(nums []int, e int) bool {
	cnt := 0
	for _, num := range nums {
		if e == num {
			cnt++
		}
	}

	return cnt > (len(nums) / 3)
}
