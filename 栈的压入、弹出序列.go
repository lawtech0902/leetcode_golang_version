package leetcode_golang_version

/*
__author__ = 'lawtech'
__date__ = '2018/8/19 下午4:43'
*/

// leetcode没有
func isPopOrder(pushV, popV []int) bool {
	var stack []int

	i := 0
	for _, v := range pushV {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popV[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return len(stack) == 0
}


