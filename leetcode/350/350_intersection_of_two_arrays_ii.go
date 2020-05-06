/*
__author__ = 'lawtech'
__date__ = '2020/05/06 1:27 下午'
*/

package _350

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	c := make(map[int]int)
	for _, v := range nums1 {
		c[v]++
	}

	var res []int
	for _, num := range nums2 {
		if val, ok := c[num]; ok && val > 0 {
			res = append(res, num)
			c[num]--
		}
	}

	return res
}
