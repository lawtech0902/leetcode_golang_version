/*
__author__ = 'lawtech'
__date__ = '2020/04/13 11:10 上午'
*/

package _238

// 左边数组乘积 * 右边数组乘积
func productExceptSelf(nums []int) []int {
	var res []int
	res = append(res, 1)

	for i := 1; i < len(nums); i++ {
		// 计算左边乘积
		res = append(res, res[i-1]*nums[i-1])
	}

	R := 1
	for i := len(nums) - 1; i >= 0; i-- {
		tmp := res[i] * R
		res[i] = tmp
		R = R * nums[i]
	}

	return res
}
