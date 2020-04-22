/*
__author__ = 'lawtech'
__date__ = '2018/8/20 上午11:46'
*/

package _315

// 暴力，也能过
func countSmaller1(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				res[i]++
			}
		}
	}

	return res
}

// 分治 + 归并
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	idxs := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idxs[i] = i
	}

	sort(nums, idxs, append([]int{}, idxs...), res, 0, len(nums)-1)
	return res
}

func sort(nums, idxs, sortedIdxs, ans []int, start, end int) {
	if start >= end {
		return
	}

	pivot := start + (end-start)/2
	sort(nums, idxs, sortedIdxs, ans, start, pivot)
	sort(nums, idxs, sortedIdxs, ans, pivot+1, end)
	idx := start
	i1, i2 := start, pivot+1
	for i1 <= pivot || i2 <= end {
		if i1 <= pivot && i2 <= end {
			if nums[idxs[i1]] > nums[idxs[i2]] {
				sortedIdxs[idx] = idxs[i2]
				i2++
			} else {
				sortedIdxs[idx] = idxs[i1]
				ans[idxs[i1]] += i2 - pivot - 1
				i1++
			}
		} else if i1 <= pivot {
			sortedIdxs[idx] = idxs[i1]
			ans[idxs[i1]] += end - pivot
			i1++
		} else if i2 <= end {
			sortedIdxs[idx] = idxs[i2]
			i2++
		}

		idx++
	}

	copy(idxs[start:end+1], sortedIdxs[start:end+1])
}
