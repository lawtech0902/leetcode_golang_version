/*
__author__ = 'lawtech'
__date__ = '2020/05/06 11:13 上午'
*/

package _349

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			if i == 0 || nums1[i] != nums1[i-1] && j == 0 || nums2[j] != nums2[j-1] {
				res = append(res, nums1[i])
			}

			i++
			j++
		}
	}

	return res
}
