package _01

/*
__author__ = 'lawtech'
__date__ = '2018/8/17 下午5:55'
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	return nil
}
