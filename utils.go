package leetcode_golang_version

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午8:22'
*/

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// min/max工具方法
func MinMax(nums []int) (int, int) {
	var max int = nums[0]
	var min int = nums[0]
	for _, value := range nums {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
