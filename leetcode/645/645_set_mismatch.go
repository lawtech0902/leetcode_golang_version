/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-03 09:49:45
 */

package _645

// hash
func findErrorNums(nums []int) []int {
	var res []int
	appear := make(map[int]bool)
	for _, num := range nums {
		if _, ok := appear[num]; ok {
			res = append(res, num)
		} else {
			appear[num] = true
		}
	}

	n := len(nums)
	for n > 0 {
		if !appear[n] {
			res = append(res, n)
		}

		n--
	}

	return res
}
