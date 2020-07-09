/*
__author__ = 'lawtech'
__date__ = '2020/07/08 4:44 下午'
*/

package _673

// dp
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	lengths := make([]int, n)
	counts := make([]int, n)
	for i := range lengths {
		lengths[i] = 1
		counts[i] = 1
	}

	maxi := -1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lengths[j]+1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					counts[i] = counts[j]
				} else if lengths[j]+1 == lengths[i] {
					counts[i] += counts[j]
				}
			}
		}

		if lengths[i] > maxi {
			maxi = lengths[i]
		}
	}

	res := 0
	for i := 0; i < len(counts); i++ {
		if lengths[i] == maxi {
			res += counts[i]
		}
	}

	return res
}
