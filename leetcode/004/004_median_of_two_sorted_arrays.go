package _04

/*
__author__ = 'lawtech'
__date__ = '2020/1/3 4:40 下午'
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOfArray(nums)
}

func combine(nums1, nums2 []int) []int {
	l1, i := len(nums1), 0
	l2, j := len(nums2), 0
	res := make([]int, l1+l2)

	for k := 0; k < l1+l2; k++ {
		if i == l1 || (i < l1 && j < l2 && nums1[i] > nums2[j]) {
			res[k] = nums2[j]
			j++
			continue
		}

		if j == l2 || (i < l1 && j < l2 && nums1[i] <= nums2[j]) {
			res[k] = nums1[i]
			i++
		}
	}

	return res
}

func medianOfArray(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}

	return float64(nums[l/2])
}
