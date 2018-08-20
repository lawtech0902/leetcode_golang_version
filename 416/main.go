package main

import (
	"fmt"
	"go_projects/leetcode_golang_version"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:57'
*/

func main() {
	nums1 := []int{1,2,3,4,6,7}
	nums2 := []int{2,3,4,8,1}

	fmt.Println(leetcode_golang_version.LengthOfLCS(nums1, nums2))
}