package _665

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:14'
*/

func checkPossibility(nums []int) bool {
	cnt := 0

	for i := 1; i < len(nums) && cnt < 2; i++ {
		if nums[i] >= nums[i-1] {
			continue
		}

		cnt++

		if i >= 2 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}

	return cnt <= 1
}
