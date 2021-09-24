package lead_to_offer_ii

func countBits(n int) []int {
	// 如果最后一位为0，则下一个数字的1个数会+1
	// 如果最后一位为1，下一个数字会导致最后一位变成0，进一位，相当于右移一位之后+1的数字中1的个数
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}

	return res
}
