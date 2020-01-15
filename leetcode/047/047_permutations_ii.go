/*
__author__ = 'lawtech'
__date__ = '2020/1/15 11:12 上午'
*/

package _47

func permuteUnique(nums []int) [][]int {
	var res [][]int
	dfs(nums, &res, 0)
	return res
}

func dfs(nums []int, res *[][]int, index int) {
	if index == len(nums) {
		*res = append(*res, dump(nums))
	}

	m := map[int]int{}
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}

		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, res, index+1)
		nums[i], nums[index] = nums[index], nums[i]
		m[nums[i]] = 1
	}
}

func dump(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
