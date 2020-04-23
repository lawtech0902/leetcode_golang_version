/*
__author__ = 'lawtech'
__date__ = '2020/04/23 10:26 上午'
*/

package _321

// 有点暴力那意思
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	maxNum := make([]int, k)
	for i := max(0, k-len(nums2)); i <= min(len(nums1), k); i++ {
		curNum := merge(choose(nums1, i), choose(nums2, k-i))
		if bigger(curNum, maxNum) {
			maxNum = curNum
		}
	}

	return maxNum
}

func choose(nums []int, k int) []int {
	choice := make([]int, k)
	leftBound := 0
LOOP:
	for i := 0; i < k; i++ {
		curMax := -1
		rightBound := len(nums) - k + i
		for j := leftBound; j <= rightBound; j++ {
			if nums[j] == 9 {
				choice[i] = 9
				leftBound = j + 1
				continue LOOP
			}

			if nums[j] > curMax {
				curMax = nums[j]
				leftBound = j + 1
			}
		}

		choice[i] = curMax
	}

	return choice
}

func merge(nums1, nums2 []int) []int {
	res := make([]int, 0)
	i1, i2 := 0, 0
	for i1 < len(nums1) && i2 < len(nums2) {
		if bigger(nums1[i1:], nums2[i2:]) {
			res = append(res, nums1[i1])
			i1++
		} else {
			res = append(res, nums2[i2])
			i2++
		}
	}

	res = append(res, nums1[i1:]...)
	res = append(res, nums2[i2:]...)
	return res
}

func bigger(nums1, nums2 []int) bool {
	minLen := min(len(nums1), len(nums2))
	for i := 0; i < minLen; i++ {
		if nums1[i] > nums2[i] {
			return true
		} else if nums1[i] < nums2[i] {
			return false
		}
	}

	// longer one is bigger
	// e.g. [2, 3] > [2], we will get 232
	//      [2, 3] < [2], we will get 223

	return len(nums1) > len(nums2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
