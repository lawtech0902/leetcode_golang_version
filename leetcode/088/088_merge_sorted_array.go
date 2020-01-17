/*
__author__ = 'lawtech'
__date__ = '2020/1/17 4:27 下午'
*/

package _88

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2[:n])
	sort.Ints(nums1)
}
