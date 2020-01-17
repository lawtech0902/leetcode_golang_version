/*
__author__ = 'lawtech'
__date__ = '2020/1/17 10:39 上午'
*/

package _75

func sortColors(nums []int) {
	// 一次扫描，3指针法
	p0, cur := 0, 0
	p2 := len(nums) - 1

	for cur <= p2 {
		if nums[cur] == 0 {
			nums[p0], nums[cur] = nums[cur], nums[p0]
			p0 += 1
			cur += 1
		} else if nums[cur] == 2 {
			nums[cur], nums[p2] = nums[p2], nums[cur]
			p2 -= 1
		} else {
			cur += 1
		}
	}
}
