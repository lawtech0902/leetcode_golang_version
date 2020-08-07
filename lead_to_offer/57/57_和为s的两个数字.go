/*
__author__ = 'lawtech'
__date__ = '2020/08/07 11:11 上午'
*/

package _57

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{num, nums[j]}
		}

		hash[num] = i
	}

	return nil
}
