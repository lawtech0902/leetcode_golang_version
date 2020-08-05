/*
__author__ = 'lawtech'
__date__ = '2020/08/05 2:46 下午'
*/

package _51

// 归并排序思想
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end-start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	var tmp []int
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}

	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}

	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}

	return cnt
}
