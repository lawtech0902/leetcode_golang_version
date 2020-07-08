/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-07 14:48:39
 */

package _667

func constructArray(n, k int) []int {
	res := make([]int, n)
	for i := 0; i < k; i++ { // 排前面k-1次落差
		if i%2 == 0 {
			res[i] = i/2 + 1 // 角标为偶数，则赋值i+1
		} else {
			res[i] = n - i/2 // 为奇数,拿n+1减去前一个
		}
	}

	// 排稳步上坡
	if (k-1)%2 == 0 {
		for i := k; i < n; i++ {
			res[i] = res[i-1] + 1
		}
	} else {
		for i := k; i < n; i++ {
			res[i] = res[i-1] - 1
		}
	}

	return res
}
