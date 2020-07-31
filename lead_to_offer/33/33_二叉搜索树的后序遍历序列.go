/*
__author__ = 'lawtech'
__date__ = '2020/07/31 4:22 下午'
*/

package _33

// 分治 递归
func verifyPostorder(postorder []int) bool {
	return helper(postorder, 0, len(postorder)-1)
}

func helper(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}

	p := i
	for postorder[p] < postorder[j] {
		p++
	}

	m := p
	for postorder[p] > postorder[j] {
		p++
	}

	return p == j && helper(postorder, i, m-1) && helper(postorder, m, j-1)
}
