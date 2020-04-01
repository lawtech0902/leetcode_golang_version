/*
__author__ = 'lawtech'
__date__ = '2020/03/30 5:16 下午'
*/

package _220

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	if t == 0 {
		m := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			if j, exist := m[nums[i]]; exist {
				if i-j <= k {
					return true
				}
			}

			m[nums[i]] = i
		}

		return false
	}

	m := make(map[int][2]int)
	for index, num := range nums {
		if num < 0 {
			num -= t - 1
		}

		if v, exist := m[num/t]; exist && index-v[0] <= k {
			return true
		}

		if v, exist := m[num/t+1]; exist && index-v[0] <= k && v[1]-num <= t {
			return true
		}

		if v, exist := m[num/t-1]; exist && index-v[0] <= k && num-v[1] <= t {
			return true
		}

		m[num/t] = [2]int{index, num}
	}

	return false
}
