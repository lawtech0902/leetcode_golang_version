/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-07 09:54:48
 */

package _658

// 双指针排除法
func findClosestElements(arr []int, k int, x int) []int {
	size := len(arr)
	left, right := 0, size-1
	removeNums := size - k
	for removeNums > 0 {
		if x-arr[left] <= arr[right]-x {
			right--
		} else {
			left++
		}

		removeNums--
	}

	return arr[left : left+k]
}
