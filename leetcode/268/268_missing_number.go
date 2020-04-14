/*
__author__ = 'lawtech'
__date__ = '2020/04/14 5:29 下午'
*/

package _268

// 数学方法，一边求和一边减
func missingNumber(nums []int) int {
	curSum := 0
	for i, num := range nums {
		curSum = curSum + num - i
	}

	return -(curSum - len(nums))
}
