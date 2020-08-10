/*
__author__ = 'lawtech'
__date__ = '2020/08/07 2:40 下午'
*/

package _59

import "container/list"

// 单调减双向队列
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	q := list.New()
	res := make([]int, len(nums)-k+1)
	for i := 0; i < k; i++ { // 未形成窗口
		for q.Len() != 0 && q.Back().Value.(int) < nums[i] {
			q.Remove(q.Back())
		}

		q.PushBack(nums[i])
	}

	res[0] = q.Front().Value.(int)
	for i := k; i < len(nums); i++ { // 形成窗口后
		if q.Front().Value.(int) == nums[i-k] {
			q.Remove(q.Front())
		}

		for q.Len() != 0 && q.Back().Value.(int) < nums[i] {
			q.Remove(q.Back())
		}

		q.PushBack(nums[i])
		res[i-k+1] = q.Front().Value.(int)
	}

	return res
}
