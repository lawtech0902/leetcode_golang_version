/*
__author__ = 'lawtech'
__date__ = '2020/06/02 10:16 上午'
*/

package _493

func reversePairs(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}

	tmp := make([]int, len(nums))
	return sortAndCount(nums, 0, len(nums)-1, tmp)
}

func merge(nums []int, low, mid, high int, tmp []int) {
	i, j, k := low, mid+1, 0
	for i <= mid && j <= high {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
			k++
		} else {
			tmp[k] = nums[j]
			j++
			k++
		}
	}

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= high {
		tmp[k] = nums[j]
		j++
		k++
	}

	copy(nums[low:high+1], tmp)
}

func count(nums []int, low, mid, high int) int {
	i, j := low, mid+1
	cnt := 0
	for i <= mid && j <= high {
		if nums[i] <= 2*nums[j] {
			i++
		} else {
			cnt += mid - i + 1
			j++
		}
	}

	return cnt
}

func sortAndCount(nums []int, low, high int, tmp []int) int {
	cnt := 0
	if low < high {
		mid := low + (high-low)/2
		cnt += sortAndCount(nums, low, mid, tmp)
		cnt += sortAndCount(nums, mid+1, high, tmp)
		cnt += count(nums, low, mid, high)
		merge(nums, low, mid, high, tmp)
	}

	return cnt
}
